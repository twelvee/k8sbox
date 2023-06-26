// Package put is used to process put commands
package put

import (
	"fmt"
	"github.com/twelvee/boxie/internal/boxie"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"strings"
)

// PutArguments is put arguments request as a struct
type PutArguments struct {
	TomlFile       string
	JsonSpec       string
	BoxType        string
	BoxName        string
	ChartLocation  string
	ValuesLocation string
	Namespace      string
}

func parseBoxFromRequest(arguments PutArguments) (structs.Box, error) {
	if len(strings.TrimSpace(arguments.TomlFile)) != 0 {
		box, err := boxie.GetTomlFormatter().GetBoxFromToml(arguments.TomlFile)
		if err != nil {
			return structs.Box{}, err
		}
		return box, nil
	}

	if len(strings.TrimSpace(arguments.JsonSpec)) != 0 {
		box, err := boxie.GetJsonFormatter().GetBoxFromJson(arguments.JsonSpec)
		if err != nil {
			return structs.Box{}, err
		}
		return box, nil
	}

	if len(strings.TrimSpace(arguments.BoxName)) == 0 || len(strings.TrimSpace(arguments.ChartLocation)) == 0 || len(strings.TrimSpace(arguments.ValuesLocation)) == 0 {
		return structs.Box{}, fmt.Errorf("Not enough arguments.")
	}

	return structs.Box{
		Name:      arguments.BoxName,
		Type:      arguments.BoxType,
		Namespace: arguments.Namespace,
		Chart:     arguments.ChartLocation,
		Values:    arguments.ValuesLocation,
	}, nil
}
