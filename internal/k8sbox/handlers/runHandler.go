package handlers

import (
	"context"
	"fmt"

	model "github.com/twelvee/k8sbox/internal/k8sbox/models"
)

func HandleRunCommand(tomlFile string, context context.Context) {
	err := model.RunEnvironment(tomlFile)
	if err != nil {
		fmt.Println("\n\rFailed to run environment. ", err)
	}
}
