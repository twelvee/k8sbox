// Package describe is used to process describe commands
package describe

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/twelvee/boxie/internal/boxie"
	"github.com/twelvee/boxie/internal/boxie/handlers"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"k8s.io/utils/strings/slices"
)

// HandleDescribeCommand is the boxie describe command handler
func HandleDescribeCommand(context context.Context, modelName string, envID string, namespace string) {
	if !slices.Contains(structs.GetEnvironmentAliases(), modelName) {
		fmt.Printf("An invalid argument. Available arguments: %s\r\n", strings.Join(structs.GetEnvironmentAliases(), ", "))
		os.Exit(1)
	}
	handlers.KuberExecutable(context, namespace)
	s := boxie.GetStorageService()
	env, err := s.GetEnvironment(namespace, envID)
	if err != nil {
		fmt.Println("Environment not found")
		os.Exit(1)
	}
	fmt.Printf("Name: %s\r\n", env.Name)
	fmt.Printf("Namespace: %s\r\n", env.Namespace)
	fmt.Println("------------------------------")
	fmt.Printf("Boxes: %d\r\n", len(env.Boxes))
	for i, b := range env.Boxes {
		fmt.Printf("Box %d (%s):\r\n", i, b.Name)
		err := boxie.GetBoxService().DescribeBoxApplications(*env, b)
		if err != nil {
			fmt.Println("Something went wrong with the box. Unable to retrieve data.")
			fmt.Println(err.Error())
			continue
		}
		fmt.Println()
	}
}
