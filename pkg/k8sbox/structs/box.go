package structs

import "helm.sh/helm/v3/pkg/release"

type Box struct {
	Type          string        `toml:"type"`
	Applications  []Application `toml:"applications"`
	Chart         string        `toml:"chart"`
	Values        string        `toml:"values"`
	Namespace     string        `toml:"namespace"`
	Name          string        `toml:"name"`
	TempDirectory string        `toml:"-"`
}

type BoxService struct {
	ProcessEnvValues   func(map[string]interface{}, string) map[string]interface{}
	ValidateBoxes      func([]Box) error
	FillEmptyFields    func(*Box, string) error
	UninstallBox       func(*Box, Environment) (*release.UninstallReleaseResponse, error)
	GetBox             func(*Box) (*release.Release, error)
	ExpandBoxVariables func([]Box) []Box
}

func Helm() string {
	return "helm"
}

func Plain() string {
	return "plain"
}
