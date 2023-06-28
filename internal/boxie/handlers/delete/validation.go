// Package delete is used to process delete commands
package delete

import (
	"fmt"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"k8s.io/utils/strings/slices"
	"strings"
)

func validateShelfDeleteRequest(resourceType string, resourceName string) error {
	if !slices.Contains(structs.GetBoxAliaces(), resourceType) {
		return fmt.Errorf("An invalid argument. Available arguments: %s\r\n", strings.Join(structs.GetBoxAliaces(), ", "))
	}
	if len(strings.TrimSpace(resourceName)) == 0 {
		return fmt.Errorf("The resource could not be deleted. None of the flags pointing to the resource are present.")
	}
	return nil
}
