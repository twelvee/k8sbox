// Package structs contain every k8sbox public structs
package structs

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// Box is your box in a struct
type Box struct {
	Type         string            `toml:"type"`
	Applications []Application     `toml:"applications"`
	Chart        string            `toml:"chart"`
	Values       string            `toml:"values"`
	Namespace    string            `toml:"namespace"`
	Name         string            `toml:"name"`
	HelmRender   map[string]string `toml:"-"`
}

// BoxService is a public BoxService
type BoxService struct {
	InstallBox              func(*Box, Environment) ([]*runtime.Object, error)
	ProcessEnvValues        func(map[string]interface{}, string) map[string]interface{}
	ValidateBoxes           func([]Box) error
	FillEmptyFields         func(Environment, *Box) error
	UninstallBox            func(Environment, Box) ([]*runtime.Object, error)
	DescribeBoxApplications func(Environment, Box) error
	ExpandBoxVariables      func([]Box) []Box
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
