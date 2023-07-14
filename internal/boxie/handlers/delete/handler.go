// Package delete is used to process delete commands
package delete

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/twelvee/boxie/internal/boxie/handlers"
	model "github.com/twelvee/boxie/internal/boxie/models"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"k8s.io/utils/strings/slices"
)

// HandleDeleteCommand is the boxie delete command handler
func HandleDeleteCommand(context context.Context, modelName string, name string, namespace string) {
	if !slices.Contains(structs.GetEnvironmentAliases(), modelName) {
		fmt.Printf("An invalid argument. Available arguments: %s\r\n", strings.Join(structs.GetEnvironmentAliases(), ", "))
		os.Exit(1)
	}
	if len(strings.TrimSpace(name)) == 0 {
		fmt.Println("The resource could not be deleted. None of the flags pointing to the resource are present.")
		os.Exit(1)
	}

	handlers.KuberExecutable(context, namespace)

	err := model.DeleteEnvironmentByName(namespace, name)
	if err != nil {
		fmt.Println("Failed to delete environment.", err)
	}
	os.Exit(1)
}
