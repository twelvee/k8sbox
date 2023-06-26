// Package delete is used to process delete commands
package delete

import (
	"context"
	"fmt"
	"github.com/twelvee/boxie/internal/boxie"
	"os"
)

// HandleShelfDeleteCommand is the boxie shelf delete command handler
func HandleShelfDeleteCommand(context context.Context, resourceType string, resourceName string) {
	err := validateShelfDeleteRequest(resourceType, resourceName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	err = shelf.DeleteBox(resourceName)
	if err != nil {
		fmt.Println("Failed to delete "+resourceType+".", err)
		os.Exit(1)
	}
	os.Exit(0)
}
