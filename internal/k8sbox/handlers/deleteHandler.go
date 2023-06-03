// Package handlers is used to process Cobra commands
package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"

	model "github.com/twelvee/k8sbox/internal/k8sbox/models"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"k8s.io/utils/strings/slices"
)

// HandleDeleteCommand is the k8sbox delete command handler
func HandleDeleteCommand(context context.Context, modelName string, tomlFile string, environmentID string) {
	if !slices.Contains(structs.GetEnvironmentAliases(), modelName) {
		fmt.Printf("An invalid argument. Available arguments: %s\r\n", strings.Join(structs.GetEnvironmentAliases(), ", "))
		os.Exit(1)
	}
	if len(strings.TrimSpace(tomlFile)) != 0 {
		err := model.DeleteEnvironmentByTomlFile(tomlFile)
		if err != nil {
			fmt.Println("Failed to delete environment.", err)
		}
		os.Exit(1)
	}
	if len(strings.TrimSpace(environmentID)) != 0 {
		err := model.DeleteEnvironmentByID(environmentID)
		if err != nil {
			fmt.Println("Failed to delete environment.", err)
		}
		os.Exit(1)
	}
	fmt.Println("The resource could not be deleted. None of the flags pointing to the resource are present.")
	os.Exit(1)
}
