// Package get is used to process get commands
package get

import (
	"fmt"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"k8s.io/utils/strings/slices"
	"strings"
)

func validateShelfGetRequest(resourceType string) error {
	if !slices.Contains(structs.GetBoxAliaces(), resourceType) {
		return fmt.Errorf("An invalid argument. Available arguments: %s", strings.Join(structs.GetBoxAliaces(), ", "))
	}
	return nil
}
