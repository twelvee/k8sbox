// Package create is used to process create commands
package create

import (
	"context"
	"fmt"
	"os"
)

// HandleCreateCommand is the boxie create command handler
func HandleCreateCommand(context context.Context, resourceType string, args CreateArguments) {
	err := validateRequest(resourceType, args)
	if err != nil {
		fmt.Println("Failed to create "+resourceType+".", err)
		os.Exit(1)
	}

	// TODO: replace with resourceType parser (currently `create box` supported only)
	_, err = parseBoxFromRequest(args)
	if err != nil {
		fmt.Println("Failed to create "+resourceType+".", err)
		os.Exit(1)
	}

	os.Exit(0)
}
