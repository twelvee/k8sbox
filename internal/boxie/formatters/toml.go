// Package formatters contains all text and file formatters
package formatters

import (
	"fmt"
	"helm.sh/helm/v3/pkg/chartutil"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	for i, b := range environment.Boxes {
		b.Chart = os.ExpandEnv(b.Chart)
		b.Values = os.ExpandEnv(b.Values)
		chartDir := b.Chart
		meta, err := os.Stat(b.Chart)
		if err != nil {
			return structs.Environment{}, fmt.Errorf("File %s not found", b.Chart)
		}
		if meta.IsDir() {
			b.Chart = b.Chart + "/" + chartutil.ChartfileName
		} else {
			chartDir = filepath.Dir(b.Chart)
		}
		data, err = os.ReadFile(b.Chart)
		if err != nil {
			return structs.Environment{}, err
		}
		environment.Boxes[i].Chart = string(data)

		if len(strings.TrimSpace(b.Values)) != 0 {
			meta, err := os.Stat(b.Values)
			if err != nil {
				return structs.Environment{}, fmt.Errorf("File %s not found", b.Values)
			}
			if meta.IsDir() {
				b.Values = b.Values + "/" + chartutil.ValuesfileName
			}
			data, err = os.ReadFile(b.Values)
			if err != nil {
				return structs.Environment{}, err
			}
			environment.Boxes[i].Values = string(data)
		}
		if environment.Boxes[i].Applications == nil {
			files, err := os.ReadDir(chartDir + "/templates")
			if err != nil {
				return environment, err
			}
			for _, f := range files {
				data, err = os.ReadFile(chartDir + "/templates/" + f.Name())
				if err != nil {
					return environment, err
				}
				environment.Boxes[i].Applications = append(environment.Boxes[i].Applications, structs.Application{Name: f.Name(), Chart: string(data)})
			}
			continue
		}
		for j, a := range environment.Boxes[i].Applications {
			data, err = os.ReadFile(os.ExpandEnv(a.Chart))
			if err != nil {
				return environment, err
			}
			environment.Boxes[i].Applications[j].Chart = string(data)
		}
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
	for i, b := range environment.Boxes {
		b.Chart = os.ExpandEnv(b.Chart)
		b.Values = os.ExpandEnv(b.Values)
		chartDir := b.Chart
		meta, err := os.Stat(b.Chart)
		if err != nil {
			return structs.Environment{}, fmt.Errorf("File %s not found", b.Chart)
		}
		if meta.IsDir() {
			b.Chart = b.Chart + "/" + chartutil.ChartfileName
		} else {
			chartDir = filepath.Dir(b.Chart)
		}
		data, err := os.ReadFile(b.Chart)
		if err != nil {
			return structs.Environment{}, err
		}
		environment.Boxes[i].Chart = string(data)

		if len(strings.TrimSpace(b.Values)) != 0 {
			meta, err := os.Stat(b.Values)
			if err != nil {
				return structs.Environment{}, fmt.Errorf("File %s not found", b.Values)
			}
			if meta.IsDir() {
				b.Values = b.Values + "/" + chartutil.ValuesfileName
			}
			data, err := os.ReadFile(b.Values)
			if err != nil {
				return structs.Environment{}, err
			}
			environment.Boxes[i].Values = string(data)
		}

		if environment.Boxes[i].Applications == nil {
			files, err := os.ReadDir(chartDir + "/templates")
			if err != nil {
				return environment, err
			}
			for _, f := range files {
				data, err = os.ReadFile(chartDir + "/templates/" + f.Name())
				if err != nil {
					return environment, err
				}
				environment.Boxes[i].Applications = append(environment.Boxes[i].Applications, structs.Application{Name: f.Name(), Chart: string(data)})
			}
			continue
		}
		for j, a := range environment.Boxes[i].Applications {
			data, err = os.ReadFile(os.ExpandEnv(a.Chart))
			if err != nil {
				return environment, err
			}
			environment.Boxes[i].Applications[j].Chart = string(data)
		}
	}
	return environment, nil
}

func getBoxFromToml(file string) (structs.Box, error) {
	var box structs.Box
	_, err := os.Stat(file)
	if err != nil {
		return structs.Box{}, fmt.Errorf("File %s not found", file)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return structs.Box{}, err
	}

	err = toml.Unmarshal(data, &box)
	if err != nil {
		return structs.Box{}, err
	}

	box.Chart = os.ExpandEnv(box.Chart)
	box.Values = os.ExpandEnv(box.Values)

	meta, err := os.Stat(box.Chart)
	if err != nil {
		return structs.Box{}, fmt.Errorf("File %s not found", box.Chart)
	}
	chartDir := box.Chart
	if meta.IsDir() {
		box.Chart = box.Chart + "/" + chartutil.ChartfileName
	} else {
		chartDir = filepath.Dir(chartDir)
	}

	meta, err = os.Stat(box.Values)
	if err != nil {
		return structs.Box{}, fmt.Errorf("File %s not found", box.Values)
	}
	if meta.IsDir() {
		box.Values = box.Values + "/" + chartutil.ValuesfileName
	}

	if len(strings.TrimSpace(box.Chart)) != 0 {
		data, err = os.ReadFile(os.ExpandEnv(box.Chart))
		if err != nil {
			return structs.Box{}, err
		}
		box.Chart = string(data)
	}

	if len(strings.TrimSpace(box.Values)) != 0 {
		data, err = os.ReadFile(os.ExpandEnv(box.Values))
		if err != nil {
			return structs.Box{}, err
		}
		box.Values = string(data)
	}

	if box.Applications == nil {
		files, err := os.ReadDir(chartDir + "/templates")
		if err != nil {
			return box, err
		}
		for _, f := range files {
			data, err = os.ReadFile(chartDir + "/templates/" + f.Name())
			if err != nil {
				return structs.Box{}, err
			}
			box.Applications = append(box.Applications, structs.Application{Name: f.Name(), Chart: string(data)})
		}
		return box, nil
	}

	for i, a := range box.Applications {
		data, err = os.ReadFile(os.ExpandEnv(a.Chart))
		if err != nil {
			return structs.Box{}, err
		}
		box.Applications[i].Chart = string(data)
	}

	return box, nil
}
