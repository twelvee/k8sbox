package services

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/k8s-box/k8sbox/pkg/k8sbox/structs"
	"github.com/k8s-box/k8sbox/pkg/k8sbox/utils"
)

func NewEnvironmentService() structs.EnvironmentService {
	return structs.EnvironmentService{
		DeployEnvironment:         deployEnvironment,
		CreateTempDeployDirectory: createTempDeployDirectory,
		ValidateEnvironment:       validateEnvironment,
	}
}

func deployEnvironment(environment structs.Environment, tempDir string) (structs.Environment, error) {
	//
	return environment, nil
}

func createTempDeployDirectory(environment structs.Environment, runDirectory string, shortID string) (string, error) {
	tempFolder, err := utils.CreateTempFolder(shortID)
	if err != nil {
		return "", err
	}

	err = moveEnvironmentFilesToTempDirectory(environment, runDirectory, tempFolder)
	if err != nil {
		return "", err
	}

	defer os.RemoveAll(tempFolder)
	return tempFolder, nil
}

func moveEnvironmentFilesToTempDirectory(environment structs.Environment, runDirectory string, tempFolder string) error {
	for _, box := range environment.Boxes {
		boxTempFolder := strings.Join([]string{tempFolder, utils.GetShortID(8)}, "/")
		os.Mkdir(boxTempFolder, 0750)
		boxChartContent, err := ioutil.ReadFile(strings.Join([]string{runDirectory, box.Chart}, ""))
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(strings.Join([]string{boxTempFolder, "Chart.yaml"}, "/"), boxChartContent, 0644)
		if err != nil {
			return err
		}

		boxValuesContent, err := ioutil.ReadFile(strings.Join([]string{runDirectory, box.Values}, ""))
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(strings.Join([]string{boxTempFolder, "Values.yaml"}, "/"), boxValuesContent, 0644)
		if err != nil {
			return err
		}

		for _, application := range box.Applications {
			applicationTempFolder := strings.Join([]string{boxTempFolder, "templates"}, "/")
			os.Mkdir(applicationTempFolder, 0750)

			applicationContent, err := ioutil.ReadFile(strings.Join([]string{runDirectory, application.Chart}, ""))
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(strings.Join([]string{applicationTempFolder, utils.GetShortID(6), ".yaml"}, ""), applicationContent, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func validateEnvironment(environment structs.Environment) error {
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
