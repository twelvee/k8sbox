// Package formatters contains all text and file formatters
package formatters

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
)

// TomlFormatter is an environment toml formatter
type TomlFormatter struct {
	GetEnvironmentFromToml func(string) (structs.Environment, error)
	GetEnvironmentViaHTTP  func(string) (structs.Environment, error)
}

// NewTomlFormatter creates a new Tomlformatter struct
func NewTomlFormatter() TomlFormatter {
	return TomlFormatter{
		GetEnvironmentFromToml: getEnvironmentFromToml,
		GetEnvironmentViaHTTP:  getEnvironmentViaHTTP,
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

func getEnvironmentViaHTTP(url string) (structs.Environment, error) {
	response, err := http.Get(url)
	var environment structs.Environment
	if err != nil {
		return environment, err
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	if err != nil {
		return environment, err
	}

	err = toml.Unmarshal(content, &environment)
	if err != nil {
		return environment, err
	}

	return environment, nil
}
