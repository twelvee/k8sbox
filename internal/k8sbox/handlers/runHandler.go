package handlers

import (
	"context"
	"fmt"
	"os"

	model "github.com/twelvee/k8sbox/internal/k8sbox/models"
)

func HandleRunCommand(context context.Context, tomlFile string) {
	err := model.RunEnvironment(tomlFile)
	if err != nil {
		fmt.Println("\n\rFailed to run environment. ", err)
		os.Exit(1)
	}
}
