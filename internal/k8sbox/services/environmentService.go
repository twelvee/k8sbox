package services

import (
	"errors"
	"strings"

	"github.com/k8s-box/k8sbox/pkg/k8sbox/structs"
)

func NewEnvironmentService() structs.EnvironmentService {
	return structs.EnvironmentService{
		DeployEnvironment:   deployEnvironment,
		ValidateEnvironment: validateEnvironment,
	}
}

func deployEnvironment(environment structs.Environment) (structs.Environment, error) {

	return environment, nil
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
