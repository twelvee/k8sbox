// Package formatters should contain toml/yaml formatters
package formatters

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
)

// Struct TomlFormatter is an environment toml formatter
type TomlFormatter struct {
	GetEnvironmentFromToml func(string) (structs.Environment, error)
}

// NewTomlFormatter creates a new Tomlformatter struct
func NewTomlFormatter() TomlFormatter {
	return TomlFormatter{
		GetEnvironmentFromToml: getEnvironmentFromToml,
	}
}

func getEnvironmentFromToml(tomlFile string) (structs.Environment, error) {
	var environment structs.Environment

	_, err := os.Stat(tomlFile)
	if err != nil {
		return structs.Environment{}, fmt.Errorf("File %s not found", tomlFile)
	}

	data, err := os.ReadFile(tomlFile)

	err = toml.Unmarshal(data, &environment)
	if err != nil {
		panic(err)
	}
	return environment, nil
}
