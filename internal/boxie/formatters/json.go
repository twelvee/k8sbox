// Package formatters contains all text and file formatters
package formatters

import (
	"encoding/json"
	"fmt"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"helm.sh/helm/v3/pkg/chartutil"
	"os"
	"path/filepath"
	"strings"
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
		data, err := os.ReadFile(os.ExpandEnv(box.Chart))
		if err != nil {
			return structs.Box{}, err
		}
		box.Chart = string(data)
	}

	if len(strings.TrimSpace(box.Values)) != 0 {
		data, err := os.ReadFile(os.ExpandEnv(box.Values))
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
			data, err := os.ReadFile(chartDir + "/templates/" + f.Name())
			if err != nil {
				return structs.Box{}, err
			}
			box.Applications = append(box.Applications, structs.Application{Name: f.Name(), Chart: string(data)})
		}
		return box, nil
	}

	for i, a := range box.Applications {
		data, err := os.ReadFile(os.ExpandEnv(a.Chart))
		if err != nil {
			return structs.Box{}, err
		}
		box.Applications[i].Chart = string(data)
	}

	return box, nil
}
