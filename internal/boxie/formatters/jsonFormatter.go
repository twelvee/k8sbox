// Package formatters contains all text and file formatters
package formatters

import (
	"encoding/json"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

// JsonFormatter is an boxie json formatter
type JsonFormatter struct {
	GetBoxFromJson func(string) (structs.Box, error)
}

// NewJsonFormatter creates a new JsonFormatter struct
func NewJsonFormatter() JsonFormatter {
	return JsonFormatter{
		GetBoxFromJson: getBoxFromJson,
	}
}

func getBoxFromJson(content string) (structs.Box, error) {
	var box structs.Box

	err := json.Unmarshal([]byte(content), &box)
	if err != nil {
		return structs.Box{}, err
	}
	return box, nil
}
