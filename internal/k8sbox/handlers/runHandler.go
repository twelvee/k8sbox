package handlers

import (
	"context"
	"fmt"

	model "github.com/k8s-box/k8sbox/internal/k8sbox/models"
	"github.com/k8s-box/k8sbox/pkg/k8sbox/structs"
)

func HandleRunCommand(tomlFile string, context context.Context) structs.Environment {
	result, err := model.RunEnvironment(tomlFile)
	if err != nil {
		panic(fmt.Sprintf("failed to run environment. %#v", err))
	}
	fmt.Println(result)
	return result
}
