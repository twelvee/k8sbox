// Package handlers is used to handle cobra commands
package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/twelvee/k8sbox/internal/k8sbox"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
)

// HandleDescribeCommand is the k8sbox describe command handler
func HandleDescribeCommand(context context.Context, getType string, envId string) {
	if getType != "environment" {
		fmt.Println("Invalid argument. Available types: environment")
		os.Exit(1)
	}
	env, err := utils.GetEnvironment(envId)
	if err != nil {
		fmt.Println("Environment not found")
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("Id: %s", env.Id))
	fmt.Println(fmt.Sprintf("Name: %s", env.Name))
	fmt.Println(fmt.Sprintf("Namespace: %s", env.Namespace))
	fmt.Println("------------------------------")
	fmt.Println(fmt.Sprintf("Boxes: %d", len(env.Boxes)))
	for i, b := range env.Boxes {
		fmt.Println(fmt.Sprintf("Box %d:", i))
		r, err := k8sbox.GetBoxService().GetBox(&b)
		if err != nil {
			fmt.Println("Something wrong with the box. Can't query.")
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(fmt.Sprintf("Name: %s", b.Name))
		fmt.Println(fmt.Sprintf("Namespace: %s", b.Namespace))
		fmt.Println("------ Chart ------")
		fmt.Println(fmt.Sprintf("API Version: %s", r.Chart.Metadata.APIVersion))
		fmt.Println(fmt.Sprintf("App Version: %s", r.Chart.Metadata.AppVersion))
		fmt.Println(fmt.Sprintf("Name: %s", r.Chart.Metadata.Name))
		fmt.Println(fmt.Sprintf("Version: %s", r.Chart.Metadata.Version))
		fmt.Println("------ Release ------")
		fmt.Println(fmt.Sprintf("Version: %d", r.Version))
		fmt.Println(fmt.Sprintf("Name: %s", r.Name))
		fmt.Println(fmt.Sprintf("Namespace: %s", r.Namespace))
		fmt.Println(fmt.Sprintf("Status: %s", r.Info.Status.String()))
		fmt.Println(fmt.Sprintf("Notes: %s", r.Info.Notes))
		fmt.Println(fmt.Sprintf("First deployed: %s", r.Info.FirstDeployed))
		fmt.Println(fmt.Sprintf("Last deployed: %s", r.Info.LastDeployed))
		fmt.Println(fmt.Sprintf("Deleted: %s", r.Info.Deleted))
		fmt.Println("------ Labels ------")
		for _, l := range r.Labels {
			fmt.Printf("%s ", l)
		}
		fmt.Println()
	}
}
