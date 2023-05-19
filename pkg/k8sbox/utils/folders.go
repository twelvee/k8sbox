package utils

import (
	"io/ioutil"
	"strings"
)

func CreateTempFolder(name string) (string, error) {
	name, err := ioutil.TempDir("/tmp", strings.Join([]string{"k8srun", name}, ""))
	if err != nil {
		return "", err
	}
	return name, nil
}
