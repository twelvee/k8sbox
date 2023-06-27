// Package formatters contains all text and file formatters
package formatters

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

// TomlFormatter is an environment toml formatter
type TomlFormatter struct {
	GetEnvironmentFromToml func(string) (structs.Environment, error)
	GetEnvironmentViaHTTP  func(string, map[string]structs.Header) (structs.Environment, error)
	GetBoxFromToml         func(string) (structs.Box, error)
}

// NewTomlFormatter creates a new Tomlformatter struct
func NewTomlFormatter() TomlFormatter {
	return TomlFormatter{
		GetEnvironmentFromToml: getEnvironmentFromToml,
		GetEnvironmentViaHTTP:  getEnvironmentViaHTTP,
		GetBoxFromToml:         getBoxFromToml,
	}
}

func getEnvironmentFromToml(tomlFile string) (structs.Environment, error) {
	var environment structs.Environment

	_, err := os.Stat(tomlFile)
	if err != nil {
		return structs.Environment{}, fmt.Errorf("File %s not found", tomlFile)
	}

	data, err := os.ReadFile(tomlFile)
	if err != nil {
		return structs.Environment{}, err
	}

	err = toml.Unmarshal(data, &environment)
	if err != nil {
		return structs.Environment{}, err
	}

	return environment, nil
}

func getEnvironmentViaHTTP(url string, headers map[string]structs.Header) (structs.Environment, error) {
	var environment structs.Environment

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return environment, err
	}
	for _, header := range headers {
		req.Header.Set(header.Name, header.Value)
	}
	res, err := client.Do(req)
	if err != nil {
		return environment, err
	}

	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		return environment, err
	}

	err = toml.Unmarshal(content, &environment)
	if err != nil {
		return environment, err
	}
	/*
		for i, _ := range environment.Boxes {
			data, err := os.ReadFile(os.ExpandEnv(environment.Boxes[i].Chart))
			if err != nil {
				return structs.Environment{}, err
			}
			environment.Boxes[i].Chart = string(data)

			data, err = os.ReadFile(os.ExpandEnv(environment.Boxes[i].Values))
			if err != nil {
				return structs.Environment{}, err
			}
			environment.Boxes[i].Values = string(data)

			for j, _ := range environment.Boxes[i].Applications {
				data, err = os.ReadFile(os.ExpandEnv(environment.Boxes[i].Applications[j].Chart))
				if err != nil {
					return structs.Environment{}, err
				}
				environment.Boxes[i].Applications[j].Chart = string(data)
			}
		}
	*/
	return environment, nil
}

func getBoxFromToml(filepath string) (structs.Box, error) {
	var box structs.Box
	_, err := os.Stat(filepath)
	if err != nil {
		return structs.Box{}, fmt.Errorf("File %s not found", filepath)
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return structs.Box{}, err
	}

	err = toml.Unmarshal(data, &box)
	if err != nil {
		return structs.Box{}, err
	}
	/*
		data, err = os.ReadFile(os.ExpandEnv(box.Chart))
		if err != nil {
			return structs.Box{}, err
		}
		box.Chart = string(data)

		data, err = os.ReadFile(os.ExpandEnv(box.Values))
		if err != nil {
			return structs.Box{}, err
		}
		box.Values = string(data)*/

	return box, nil
}
