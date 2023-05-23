package handlers

import (
	"context"
	"fmt"
	"os"

	model "github.com/twelvee/k8sbox/internal/k8sbox/models"
)

func HandleDeleteCommand(tomlFile string, context context.Context) {
	err := model.DeleteEnvironment(tomlFile)
	if err != nil {
		fmt.Println("\n\rFailed to delete environment. ", err)
		os.Exit(1)
	}
}
