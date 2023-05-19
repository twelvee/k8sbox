package services

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/k8s-box/k8sbox/pkg/k8sbox/structs"
)

func NewBoxService() structs.BoxService {
	return structs.BoxService{
		ValidateBoxes: validateBoxes,
	}
}

func validateBoxes(boxes []structs.Box, runDirectory string, applicationService structs.ApplicationService) error {
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

		applicationsErrors := applicationService.ValidateApplications(box.Applications, runDirectory)
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
