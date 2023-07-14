// Package structs contain every boxie public structs
package structs

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// Box is your box in a struct
type Box struct {
	Type              string            `toml:"type"`
	Applications      []Application     `toml:"applications"`
	Chart             string            `toml:"chart"`
	Values            string            `toml:"values"`
	VariablesLocation string            `toml:"variables"`
	Variables         map[string]string `toml:"-"`
	Namespace         string            `toml:"namespace"`
	Name              string            `toml:"name"`
	HelmRender        map[string]string `toml:"-"`
	Created           string            `toml:"-"`
}

// BoxService is a public BoxService
type BoxService struct {
	InstallBox                  func(*Box, Environment) ([]*runtime.Object, error)
	ProcessEnvironmentEnvValues func(map[string]interface{}, Environment) map[string]interface{}
	ProcessBoxEnvValues         func(map[string]interface{}, Box) map[string]interface{}
	ValidateBoxes               func([]Box) error
	FillEmptyFields             func(Environment, *Box) error
	UninstallBox                func(Environment, Box) ([]*runtime.Object, error)
	DescribeBoxApplications     func(Environment, Box) error
	ExpandBoxVariables          func([]Box) []Box
}

// Helm is helm string getter
func Helm() string {
	return "helm"
}

// Plain is plain string getter
func Plain() string {
	return "plain"
}

// GetBoxAliaces return a slice of box model name aliases
func GetBoxAliaces() []string {
	return []string{"box", "boxes"}
}

// DeleteBoxRequest is rest api request that delete box by its name
type DeleteBoxRequest struct {
	Name string
}

// GetBoxRequest is rest api request that gets box by its name
type GetBoxRequest struct {
	Name string
}
