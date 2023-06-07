// Package handlers is used to process Cobra commands
package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/twelvee/k8sbox/internal/k8sbox"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
	"k8s.io/utils/strings/slices"
)

// HandleDescribeCommand is the k8sbox describe command handler
func HandleDescribeCommand(context context.Context, modelName string, envID string) {
	if !slices.Contains(structs.GetEnvironmentAliases(), modelName) {
		fmt.Printf("An invalid argument. Available arguments: %s\r\n", strings.Join(structs.GetEnvironmentAliases(), ", "))
		os.Exit(1)
	}
	env, err := utils.GetEnvironment(envID)
	if err != nil {
		fmt.Println("Environment not found")
		os.Exit(1)
	}
	fmt.Printf("Id: %s\r\n", env.ID)
	fmt.Printf("Name: %s\r\n", env.Name)
	fmt.Printf("Namespace: %s\r\n", env.Namespace)
	fmt.Println("------------------------------")
	fmt.Printf("Boxes: %d\r\n", len(env.Boxes))
	for i, b := range env.Boxes {
		fmt.Printf("Box %d (%s):\r\n", i, b.Name)
		err := k8sbox.GetBoxService().DescribeBoxApplications(&b, *env)
		if err != nil {
			fmt.Println("Something went wrong with the box. Unable to retrieve data.")
			fmt.Println(err.Error())
			continue
		}
		fmt.Println()
	}
}
