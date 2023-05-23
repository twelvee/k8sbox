package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
)

func NewApplicationService() structs.ApplicationService {
	return structs.ApplicationService{
		ValidateApplications: validateApplications,
		ExpandApplications:   ExpandApplications,
	}
}

func validateApplications(applications []structs.Application) []string {
	var messages []string
	for index, application := range applications {
		if len(application.Name) == 0 {
			messages = append(messages, fmt.Sprintf("--> Application %d: Name is missing", index))
		}

		if len(strings.TrimSpace(application.Chart)) == 0 {
			messages = append(messages, fmt.Sprintf("--> Application %d: Chart is missing", index))
		}
		_, err := os.Stat(application.Chart)
		if err != nil {
			messages = append(messages, fmt.Sprintf("--> Application %d: Chart file can't be opened (%s)", index, application.Chart))
		}
	}
	return messages
}

func ExpandApplications(applications []structs.Application) []structs.Application {
	var newApplications []structs.Application
	for _, a := range applications {
		a.Name = os.ExpandEnv(a.Name)
		a.Chart = os.ExpandEnv(a.Chart)
		newApplications = append(newApplications, a)
	}
	return newApplications
}
