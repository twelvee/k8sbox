package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/release"
)

func NewBoxService() structs.BoxService {
	return structs.BoxService{
		ValidateBoxes:   validateBoxes,
		FillEmptyFields: fillEmptyFields,
		UninstallBox:    UninstallBox,
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

func validateBoxes(boxes []structs.Box, runDirectory string) error {
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
		chartFilePath := strings.Join([]string{runDirectory, box.Chart}, "")
		_, err := os.Stat(chartFilePath)
		if err != nil {
			messages = append(messages, fmt.Sprintf("-> Box %d: Chart file can't be opened (%s)", index, chartFilePath))
		}

		if len(strings.TrimSpace(box.Values)) == 0 {
			messages = append(messages, fmt.Sprintf("-> Box %d: Values are missing", index))
		}
		valuesFilePath := strings.Join([]string{runDirectory, box.Values}, "")
		_, err = os.Stat(valuesFilePath)
		if err != nil {
			messages = append(messages, fmt.Sprintf("-> Box %d: Values file can't be opened (%s)", index, valuesFilePath))
		}

		applicationsErrors := validateApplications(box.Applications, runDirectory)
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

func InstallBox(box *structs.Box, environmentId string) (*release.Release, error) {
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
	r, err := client.RunWithContext(context.Background(), chart, chart.Values)
	if err != nil {
		return r, err
	}
	utils.SaveBox(*box, environmentId)
	return r, nil
}

func UpgradeBox(box *structs.Box, environmentId string) (*release.Release, error) {
	config := GetActionConfig(box.Namespace)
	client := action.NewUpgrade(config)
	client.Namespace = box.Namespace
	client.Install = true
	client.CleanupOnFail = true
	chart, err := loader.Load(box.TempDirectory)
	if err != nil {
		return nil, err
	}
	r, err := client.RunWithContext(context.Background(), box.Name, chart, chart.Values)
	if err != nil {
		return r, err
	}
	utils.SaveBox(*box, environmentId)
	return r, nil
}

func UninstallBox(box *structs.Box, environmentId string) (*release.UninstallReleaseResponse, error) {
	config := GetActionConfig(box.Namespace)
	client := action.NewUninstall(config)
	r, err := client.Run(box.Name)
	if err != nil {
		return r, err
	}
	err = utils.RemoveBox(*box, environmentId)
	if err != nil {
		return r, err
	}
	return r, nil
}

func GetBox(box structs.Box) (*release.Release, error) {
	return nil, nil
}
