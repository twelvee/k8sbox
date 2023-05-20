package handlers

import (
	"context"
	"fmt"

	model "github.com/twelvee/k8sbox/internal/k8sbox/models"
)

func HandleRunCommand(tomlFile string, context context.Context) {
	err := model.RunEnvironment(tomlFile)
	if err != nil {
		panic(fmt.Sprintf("failed to run environment. %#v", err))
	}
}
