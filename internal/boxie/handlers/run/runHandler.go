// Package run is used to process run command
package run

import (
	"context"
	"fmt"
	"os"

	model "github.com/twelvee/boxie/internal/boxie/models"
)

// HandleRunCommand is the boxie run command handler
func HandleRunCommand(context context.Context, tomlFile string) {
	err := model.RunEnvironment(tomlFile)
	if err != nil {
		fmt.Println("Failed to run environment. ", err)
		os.Exit(1)
	}
}
