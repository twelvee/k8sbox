// Package handlers is used to handle cobra commands
package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"

	model "github.com/twelvee/k8sbox/internal/k8sbox/models"
)

// HandleDeleteCommand is the k8sbox delete command handler
func HandleDeleteCommand(context context.Context, tomlFile string, environmentID string) {
	if len(strings.TrimSpace(tomlFile)) != 0 {
		err := model.DeleteEnvironmentByTomlFile(tomlFile)
		if err != nil {
			fmt.Println("\n\rFailed to delete environment.", err)
		}
		os.Exit(1)
	}
	if len(strings.TrimSpace(environmentID)) != 0 {
		err := model.DeleteEnvironmentByID(environmentID)
		if err != nil {
			fmt.Println("\n\rFailed to delete environment.", err)
		}
		os.Exit(1)
	}
	fmt.Println("\n\rFailed to delete environment. No environment ID or environment toml file present.")
	os.Exit(1)
}
