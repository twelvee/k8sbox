package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
)

func NewEnvironmentService() structs.EnvironmentService {
	return structs.EnvironmentService{
		DeployEnvironment:         deployEnvironment,
		CreateTempDeployDirectory: createTempDeployDirectory,
		ValidateEnvironment:       validateEnvironment,
	}
}

func GetActionConfig(namespace string) *action.Configuration {
	restClientGetter := kube.GetConfig(os.Getenv("KUBECONFIG"), "minikube", namespace)
	actionConfig := new(action.Configuration)
	actionConfig.Init(restClientGetter, namespace, "secret", func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	})

	return actionConfig
}

func deployEnvironment(environment *structs.Environment) error {
	for _, box := range environment.Boxes {
		UninstallBox(&box)

		_, err := InstallBox(&box)
		if err != nil {
			return err
		}
	}

	//return releases, nil
	return nil
}

func createTempDeployDirectory(environment *structs.Environment, runDirectory string, shortID string) (string, error) {
	tempFolder, err := utils.CreateTempFolder(shortID)
	if err != nil {
		return "", err
	}
	environment.TempDirectory = tempFolder
	err = moveEnvironmentFilesToTempDirectory(environment, runDirectory)
	if err != nil {
		return "", err
	}
	return tempFolder, nil
}

func moveEnvironmentFilesToTempDirectory(environment *structs.Environment, runDirectory string) error {
	for bi, _ := range environment.Boxes {
		environment.Boxes[bi].TempDirectory = strings.Join([]string{environment.TempDirectory, utils.GetShortID(8)}, "/")
		os.Mkdir(environment.Boxes[bi].TempDirectory, 0750)
		boxChartContent, err := ioutil.ReadFile(strings.Join([]string{runDirectory, environment.Boxes[bi].Chart}, ""))
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(strings.Join([]string{environment.Boxes[bi].TempDirectory, "Chart.yaml"}, "/"), boxChartContent, 0644)
		if err != nil {
			return err
		}

		boxValuesContent, err := ioutil.ReadFile(strings.Join([]string{runDirectory, environment.Boxes[bi].Values}, ""))
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(strings.Join([]string{environment.Boxes[bi].TempDirectory, "values.yaml"}, "/"), boxValuesContent, 0644)
		if err != nil {
			return err
		}

		for ai, _ := range environment.Boxes[bi].Applications {
			environment.Boxes[bi].Applications[ai].TempDirectory = strings.Join([]string{environment.Boxes[bi].TempDirectory, "templates"}, "/")
			os.Mkdir(environment.Boxes[bi].Applications[ai].TempDirectory, 0750)

			applicationContent, err := ioutil.ReadFile(strings.Join([]string{runDirectory, environment.Boxes[bi].Applications[ai].Chart}, ""))
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(strings.Join([]string{environment.Boxes[bi].Applications[ai].TempDirectory, "/", utils.GetShortID(6), ".yaml"}, ""), applicationContent, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateEnvironment(environment *structs.Environment) error {
	var messages []string
	if len(strings.TrimSpace(environment.Name)) == 0 {
		messages = append(messages, "Environment name is missing")
	}

	if len(environment.Boxes) == 0 {
		messages = append(messages, "Environment boxes are missing")
	}

	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n\r"))
	}

	return nil
}
