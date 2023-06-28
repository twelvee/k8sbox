// Package run is used to process run command
package run

import (
	"fmt"
	"strings"
)

func validateRequest(arguments RunArguments) error {
	if len(strings.TrimSpace(arguments.Name)) == 0 || len(strings.TrimSpace(arguments.Namespace)) == 0 || arguments.Boxes == nil {
		return fmt.Errorf("Not enough arguments. Name, namespace and boxes are required.")
	}

	return nil
}
