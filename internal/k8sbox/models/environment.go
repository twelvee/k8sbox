package model

import (
	"github.com/k8s-box/k8sbox/internal/k8sbox"
	"github.com/k8s-box/k8sbox/pkg/k8sbox/structs"
)

func RunEnvironment(tomlFile string) (structs.Environment, error) {
	environment := k8sbox.GetTomlFormatter().GetEnvironmentFromToml(tomlFile)
	result, err := k8sbox.GetEnvironmentService().DeployEnvironment(environment)
	if err != nil {
		panic(err)
	}

	return result, nil
}
