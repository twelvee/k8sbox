package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	model "github.com/twelvee/k8sbox/internal/k8sbox/models"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
)

func HandleRunCommand(tomlFile string, context context.Context) structs.Environment {
	result, err := model.RunEnvironment(tomlFile)
	if err != nil {
		panic(fmt.Sprintf("failed to run environment. %#v", err))
	}
	s, _ := json.Marshal(result)
	fmt.Println(string(s))
	return result
}
