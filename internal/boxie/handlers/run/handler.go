// Package run is used to process run command
package run

import (
	"context"
	"fmt"
	"github.com/twelvee/boxie/internal/boxie"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"github.com/twelvee/boxie/pkg/boxie/utils"
	"os"
	"strings"

	model "github.com/twelvee/boxie/internal/boxie/models"
)

// HandleRunCommand is the boxie run command handler
func HandleRunCommand(context context.Context, args RunArguments) {
	if len(strings.TrimSpace(args.TomlFile)) != 0 {
		environment, err := boxie.GetTomlFormatter().GetEnvironmentFromToml(args.TomlFile)
		if err != nil {
			fmt.Println("Failed to run environment. ", err)
			os.Exit(1)
		}
		err = model.RunEnvironment(environment)
		if err != nil {
			fmt.Println("Failed to run environment. ", err)
			os.Exit(1)
		}
		return
	}

	err := validateRequest(args)
	if err != nil {
		fmt.Println("Failed to run environment. ", err)
		os.Exit(1)
	}

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))

	var boxes []structs.Box
	for _, v := range args.Boxes {
		b, err := shelf.GetBox(v)
		if err != nil {
			fmt.Println("Failed to run environment. ", err)
			os.Exit(1)
		}

		boxes = append(boxes, b)
	}

	envID := args.ID
	if len(strings.TrimSpace(envID)) == 0 {
		envID = utils.GetShortID(6)
	}

	environment := structs.Environment{
		ID:           envID,
		Name:         args.Name,
		Namespace:    args.Namespace,
		Boxes:        boxes,
		Variables:    args.VariablesPath,
		VariablesMap: args.Variables,
	}

	err = model.RunEnvironment(environment)
	if err != nil {
		fmt.Println("Failed to run environment. ", err)
		os.Exit(1)
	}
}
