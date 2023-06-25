// Package create is used to process create commands
package create

import (
	"fmt"
	"github.com/twelvee/boxie/internal/boxie"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"k8s.io/utils/strings/slices"
	"strings"
)

func validateRequest(resourceType string, arguments CreateArguments) error {
	if !slices.Contains(structs.GetBoxAliaces(), resourceType) {
		return fmt.Errorf("An invalid argument. Available arguments: %s\r\n", strings.Join(structs.GetBoxAliaces(), ", "))
	}

	if len(strings.TrimSpace(arguments.TomlFile)) != 0 && len(strings.TrimSpace(arguments.JsonSpec)) != 0 {
		return fmt.Errorf("Json specification and toml file location specified at the same time.")
	}

	if len(strings.TrimSpace(arguments.TomlFile)) != 0 {
		_, err := boxie.GetTomlFormatter().GetBoxFromToml(arguments.TomlFile)
		if err != nil {
			return err
		}
		return nil
	}

	if len(strings.TrimSpace(arguments.JsonSpec)) != 0 {
		_, err := boxie.GetJsonFormatter().GetBoxFromJson(arguments.JsonSpec)
		if err != nil {
			return err
		}
		return nil
	}

	if len(strings.TrimSpace(arguments.BoxName)) == 0 || len(strings.TrimSpace(arguments.ChartLocation)) == 0 || len(strings.TrimSpace(arguments.ValuesLocation)) == 0 {
		return fmt.Errorf("Not enough arguments.")
	}

	return nil
}
