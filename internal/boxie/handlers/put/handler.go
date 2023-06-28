// Package put is used to process put commands
package put

import (
	"context"
	"fmt"
	"github.com/twelvee/boxie/internal/boxie"
	"os"
)

// HandlePutCommand is the boxie put command handler
func HandlePutCommand(context context.Context, resourceType string, args PutArguments, force bool) {
	err := validateRequest(resourceType, args)
	if err != nil {
		fmt.Println("Failed to put "+resourceType+" into shelf.", err)
		os.Exit(1)
	}

	// TODO: replace with resourceType parser (currently `put box` supported only)
	box, err := parseBoxFromRequest(args)
	if err != nil {
		fmt.Println("Failed to put "+resourceType+" into shelf.", err)
		os.Exit(1)
	}

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	err = shelf.PutBox(box, force)
	if err != nil {
		fmt.Println("Failed to put "+resourceType+" into shelf.", err)
		os.Exit(1)
	}

	os.Exit(0)
}
