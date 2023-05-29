// Package structs contain every k8sbox public structs
package structs

import "helm.sh/helm/v3/pkg/release"

// Box is your box in a struct
type Box struct {
	Type          string        `toml:"type"`
	Applications  []Application `toml:"applications"`
	Chart         string        `toml:"chart"`
	Values        string        `toml:"values"`
	Namespace     string        `toml:"namespace"`
	Name          string        `toml:"name"`
	TempDirectory string        `toml:"-"`
}

// BoxService is a public BoxService
type BoxService struct {
	ProcessEnvValues   func(map[string]interface{}, string) map[string]interface{}
	ValidateBoxes      func([]Box) error
	FillEmptyFields    func(*Box, string) error
	UninstallBox       func(*Box, Environment) (*release.UninstallReleaseResponse, error)
	GetBox             func(*Box) (*release.Release, error)
	ExpandBoxVariables func([]Box) []Box
}

// Helm is helm string getter
func Helm() string {
	return "helm"
}

// Plain is plain string getter
func Plain() string {
	return "plain"
}
