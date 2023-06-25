// Package handlers is used to process Cobra commands
package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"

	model "github.com/twelvee/boxie/internal/boxie/models"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"k8s.io/utils/strings/slices"
)

// HandleDeleteCommand is the boxie delete command handler
func HandleDeleteCommand(context context.Context, modelName string, environmentID string, namespace string) {
	if !slices.Contains(structs.GetEnvironmentAliases(), modelName) {
		fmt.Printf("An invalid argument. Available arguments: %s\r\n", strings.Join(structs.GetEnvironmentAliases(), ", "))
		os.Exit(1)
	}
	if len(strings.TrimSpace(environmentID)) == 0 {
		fmt.Println("The resource could not be deleted. None of the flags pointing to the resource are present.")
		os.Exit(1)
	}

	KuberExecutable(context, namespace)

	err := model.DeleteEnvironmentByID(namespace, environmentID)
	if err != nil {
		fmt.Println("Failed to delete environment.", err)
	}
	os.Exit(1)
}
