// Package services contains buisness-logic methods of the models
package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/release"
)

// NewBoxService creates a new BoxService
func NewBoxService() structs.BoxService {
	return structs.BoxService{
		ProcessEnvValues:   processEnvValues,
		ValidateBoxes:      validateBoxes,
		FillEmptyFields:    fillEmptyFields,
		UninstallBox:       UninstallBox,
		GetBox:             GetBox,
		ExpandBoxVariables: expandBoxVariables,
	}
}

func fillEmptyFields(box *structs.Box, environmentNamespace string) error {
	if len(strings.TrimSpace(box.Namespace)) == 0 {
		if len(strings.TrimSpace(environmentNamespace)) == 0 {
			box.Namespace = strings.ToLower(strings.Join([]string{"k8srun", utils.GetShortNamespace(8)}, "-"))
		} else {
			box.Namespace = environmentNamespace
		}
	}
	if len(strings.TrimSpace(box.Name)) == 0 {
		box.Name = strings.ToLower(strings.Join([]string{"k8srun", utils.GetShortNamespace(8)}, "-"))
	}

	return nil
}

func processEnvValues(values map[string]interface{}, dotenvPath string) map[string]interface{} {
	if len(dotenvPath) > 0 {
		godotenv.Load(dotenvPath)
	}
	for k, v := range values {
		if len(os.ExpandEnv(fmt.Sprintf("%v", v))) == 0 {
			continue
		}
		values[k] = os.ExpandEnv(fmt.Sprintf("%v", v))
	}
	return values
}

func validateBoxes(boxes []structs.Box) error {
	var messages []string
	for index, box := range boxes {
		if len(strings.TrimSpace(box.Type)) == 0 {
			messages = append(messages, fmt.Sprintf("-> Box %d: Type is missing", index))
		}

		if len(box.Applications) == 0 {
			messages = append(messages, fmt.Sprintf("-> Box %d: Applications are missing", index))
		}

		if len(strings.TrimSpace(box.Chart)) == 0 {
			messages = append(messages, fmt.Sprintf("-> Box %d: Chart is missing", index))
		}
		_, err := os.Stat(box.Chart)
		if err != nil {
			messages = append(messages, fmt.Sprintf("-> Box %d: Chart file can't be opened (%s)", index, box.Chart))
		}

		if len(strings.TrimSpace(box.Values)) == 0 {
			messages = append(messages, fmt.Sprintf("-> Box %d: Values are missing", index))
		}
		_, err = os.Stat(box.Values)
		if err != nil {
			messages = append(messages, fmt.Sprintf("-> Box %d: Values file can't be opened (%s)", index, box.Values))
		}

		applicationsErrors := validateApplications(box.Applications)
		if len(applicationsErrors) > 0 {
			for _, err := range applicationsErrors {
				messages = append(messages, fmt.Sprintf("-> Box %d: \n\r%s", index, err))
			}
		}
	}
	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n\r"))
	}
	return nil
}

// InstallBox will deploy your box applications into your k8s cluster
func InstallBox(box *structs.Box, environment structs.Environment) (*release.Release, error) {
	config := GetActionConfig(box.Namespace)
	client := action.NewInstall(config)
	client.UseReleaseName = true
	client.Namespace = box.Namespace
	client.ReleaseName = box.Name
	client.CreateNamespace = true
	client.Replace = true
	chart, err := loader.Load(box.TempDirectory)
	if err != nil {
		return nil, err
	}
	r, err := client.RunWithContext(context.Background(), chart, processEnvValues(chart.Values, environment.Variables))
	if err != nil {
		return r, err
	}
	utils.SaveBox(*box, environment.Id)
	return r, nil
}

// UpgradeBox will upgrade your box applications in your k8s cluster
func UpgradeBox(box *structs.Box, environment structs.Environment) (*release.Release, error) {
	config := GetActionConfig(box.Namespace)
	client := action.NewUpgrade(config)
	client.Namespace = box.Namespace
	client.Install = true
	client.CleanupOnFail = true
	chart, err := loader.Load(box.TempDirectory)
	if err != nil {
		return nil, err
	}

	r, err := client.RunWithContext(context.Background(), box.Name, chart, processEnvValues(chart.Values, environment.Variables))
	if err != nil {
		return r, err
	}
	utils.SaveBox(*box, environment.Id)
	return r, nil
}

// Uninstall will uninstall your box applications from your k8s cluster
func UninstallBox(box *structs.Box, environment structs.Environment) (*release.UninstallReleaseResponse, error) {
	config := GetActionConfig(box.Namespace)
	client := action.NewUninstall(config)
	r, err := client.Run(box.Name)
	if err != nil {
		return r, err
	}
	err = utils.RemoveBox(*box, environment.Id)
	if err != nil {
		return r, err
	}
	return r, nil
}

// GetBox will return a helm release from your k8s cluster
func GetBox(box *structs.Box) (*release.Release, error) {
	config := GetActionConfig(box.Namespace)
	client := action.NewGet(config)

	r, err := client.Run(box.Name)
	return r, err
}

func expandBoxVariables(boxes []structs.Box) []structs.Box {
	var newBoxes []structs.Box

	for _, b := range boxes {
		b.Name = os.ExpandEnv(b.Name)
		b.Namespace = os.ExpandEnv(b.Namespace)
		b.Type = os.ExpandEnv(b.Type)
		b.Chart = os.ExpandEnv(b.Chart)
		b.Values = os.ExpandEnv(b.Values)
		b.Applications = ExpandApplications(b.Applications)
		newBoxes = append(newBoxes, b)
	}
	return newBoxes
}
