// Package handlers is used to process Cobra commands
package handlers

import (
	"context"
	"fmt"
	"os"

	model "github.com/twelvee/k8sbox/internal/k8sbox/models"
)

// HandleRunCommand is the k8sbox run command handler
func HandleRunCommand(context context.Context, tomlFile string) {
	err := model.RunEnvironment(tomlFile)
	if err != nil {
		fmt.Println("Failed to run environment. ", err)
		os.Exit(1)
	}
}
