package formatters

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
)

type TomlFormatter struct {
	GetEnvironmentFromToml func(string) (structs.Environment, error)
}

func NewTomlFormatter() TomlFormatter {
	return TomlFormatter{
		GetEnvironmentFromToml: getEnvironmentFromToml,
	}
}

func getEnvironmentFromToml(tomlFile string) (structs.Environment, error) {
	var environment structs.Environment

	_, err := os.Stat(tomlFile)
	if err != nil {
		return structs.Environment{}, errors.New(fmt.Sprintf("File %s not found", tomlFile))
	}

	data, err := os.ReadFile(tomlFile)

	err = toml.Unmarshal(data, &environment)
	if err != nil {
		panic(err)
	}
	return environment, nil
}
