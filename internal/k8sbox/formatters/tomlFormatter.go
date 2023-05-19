package formatters

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/k8s-box/k8sbox/pkg/k8sbox/structs"
)

type TomlFormatter struct {
	GetEnvironmentFromToml func(string) (structs.Environment, string)
}

func NewTomlFormatter() TomlFormatter {
	return TomlFormatter{
		GetEnvironmentFromToml: getEnvironmentFromToml,
	}
}

func getEnvironmentFromToml(tomlFile string) (structs.Environment, string) {
	var environment structs.Environment
	data, err := os.ReadFile(tomlFile)
	info, err := os.Stat(tomlFile)
	boxesPath := strings.ReplaceAll(tomlFile, info.Name(), "")
	if err != nil {
		panic(fmt.Sprintf("File %s not found", tomlFile))
	}
	err = toml.Unmarshal(data, &environment)
	if err != nil {
		panic(err)
	}
	return environment, boxesPath
}
