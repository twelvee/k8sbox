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
		fmt.Printf("Box %d:\r\n", i)
		r, err := k8sbox.GetBoxService().GetBox(&b)
		if err != nil {
			fmt.Println("Something went wrong with the box. Unable to retrieve data.")
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("Name: %s\r\n", b.Name)
		fmt.Printf("Namespace: %s\r\n", b.Namespace)
		fmt.Println("------ Chart ------")
		fmt.Printf("API Version: %s\r\n", r.Chart.Metadata.APIVersion)
		fmt.Printf("App Version: %s\r\n", r.Chart.Metadata.AppVersion)
		fmt.Printf("Name: %s\r\n", r.Chart.Metadata.Name)
		fmt.Printf("Version: %s\r\n", r.Chart.Metadata.Version)
		fmt.Println("------ Release ------")
		fmt.Printf("Version: %d\r\n", r.Version)
		fmt.Printf("Name: %s\r\n", r.Name)
		fmt.Printf("Namespace: %s\r\n", r.Namespace)
		fmt.Printf("Status: %s\r\n", r.Info.Status.String())
		fmt.Printf("Notes: %s\r\n", r.Info.Notes)
		fmt.Printf("First deployed: %s\r\n", r.Info.FirstDeployed)
		fmt.Printf("Last deployed: %s\r\n", r.Info.LastDeployed)
		fmt.Printf("Deleted: %s\r\n", r.Info.Deleted)
		fmt.Printf("------ Labels ------")
		for _, l := range r.Labels {
			fmt.Printf("%s ", l)
		}
		fmt.Println()
	}
}
