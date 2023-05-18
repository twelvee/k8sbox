package services

import "github.com/k8s-box/k8sbox/pkg/k8sbox/structs"

func NewEnvironmentService() structs.EnvironmentService {
	return structs.EnvironmentService{
		DeployEnvironment: deployEnvironment,
	}
}

func deployEnvironment(environment structs.Environment) (structs.Environment, error) {
	return environment, nil
}
