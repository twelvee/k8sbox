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
		DeleteEnvironment:         deleteEnvironment,
		CreateTempDeployDirectory: createTempDeployDirectory,
		ValidateEnvironment:       validateEnvironment,
		ExpandVariables:           expandVariables,
	}
}

func expandVariables(environment *structs.Environment) {
	environment.Name = os.ExpandEnv(environment.Name)
	environment.Id = os.ExpandEnv(environment.Id)
	environment.Namespace = os.ExpandEnv(environment.Namespace)
	environment.Variables = os.ExpandEnv(environment.Variables)
}

func deleteEnvironment(environment *structs.Environment) error {
	err := utils.RemoveEnvironment(environment.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetActionConfig(namespace string) *action.Configuration {
	restClientGetter := kube.GetConfig(os.Getenv("KUBECONFIG"), "minikube", namespace)
	actionConfig := new(action.Configuration)
	actionConfig.Init(restClientGetter, namespace, "secret", func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	})

	return actionConfig
}

func deployEnvironment(environment *structs.Environment, isSaved bool) error {
	if !isSaved {
		utils.SaveEnvironment(*environment)
	}

	for _, box := range environment.Boxes {
		_, err := InstallBox(&box, *environment)
		if err != nil {
			return err
		}
	}

	//return releases, nil
	return nil
}

func createTempDeployDirectory(environment *structs.Environment, runDirectory string, isSavedAlready bool) (string, error) {
	if isSavedAlready {
		env, err := utils.GetEnvironment(environment.Id)
		if err != nil {
			return "", err
		}
		environment.TempDirectory = env.TempDirectory
		err = moveEnvironmentFilesToTempDirectory(environment, runDirectory)
		if err != nil {
			return "", err
		}
		return env.TempDirectory, nil
	}
	tempFolder, err := utils.CreateTempFolder(utils.GetShortID(8))
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
	envVariablesContent, err := ioutil.ReadFile(strings.Join([]string{runDirectory, environment.Variables}, ""))
	if err != nil {
		return err
	}

	environment.Variables = strings.Join([]string{environment.TempDirectory, ".env"}, "/")
	err = ioutil.WriteFile(environment.Variables, envVariablesContent, 0644)
	if err != nil {
		return err
	}
	for bi, box := range environment.Boxes {
		saved, err := utils.IsBoxSaved(environment.Id, box)
		if err != nil {
			return err
		}
		if saved {
			return nil
		}

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

func validateEnvironment(environment *structs.Environment, runDirectory string) error {
	var messages []string
	if len(strings.TrimSpace(environment.Id)) == 0 {
		messages = append(messages, "Environment id is missing")
	}

	if len(strings.TrimSpace(environment.Name)) == 0 {
		messages = append(messages, "Environment name is missing")
	}

	if len(strings.TrimSpace(environment.Variables)) > 0 {
		envFilePath := strings.Join([]string{runDirectory, environment.Variables}, "")
		_, err := os.Stat(envFilePath)
		if err != nil {
			messages = append(messages, fmt.Sprintf("Environment variables specified but file missing (%s)", envFilePath))
		}
	}

	if len(environment.Boxes) == 0 {
		messages = append(messages, "Environment boxes are missing")
	}

	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n\r"))
	}

	return nil
}
