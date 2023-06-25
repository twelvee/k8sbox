// Package utils is a useful utils that boxie use. Methods are usually exported
package utils

import (
	"io/ioutil"
	"strings"
)

// CreateTempFolder creates a temp directory for an environment
func CreateTempFolder(name string) (string, error) {
	name, err := ioutil.TempDir("/tmp", strings.Join([]string{"k8srun", name}, ""))
	if err != nil {
		return "", err
	}
	return name, nil
}
