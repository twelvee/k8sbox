package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/k8s-box/k8sbox/pkg/k8sbox/structs"
)

func NewApplicationService() structs.ApplicationService {
	return structs.ApplicationService{
		ValidateApplications: validateApplications,
	}
}

func validateApplications(applications []structs.Application, runDirectory string) []string {
	var messages []string
	for index, application := range applications {
		if len(application.Name) == 0 {
			messages = append(messages, fmt.Sprintf("--> Application %d: Name is missing", index))
		}

		if len(strings.TrimSpace(application.Chart)) == 0 {
			messages = append(messages, fmt.Sprintf("--> Application %d: Chart is missing", index))
		}
		chartFilePath := strings.Join([]string{runDirectory, application.Chart}, "")
		_, err := os.Stat(chartFilePath)
		if err != nil {
			messages = append(messages, fmt.Sprintf("--> Application %d: Chart file can't be opened (%s)", index, chartFilePath))
		}
	}
	return messages
}
