package structs

import "helm.sh/helm/v3/pkg/release"

type Box struct {
	Type          string
	Applications  []Application
	Chart         string
	Values        string
	Namespace     string
	Name          string
	TempDirectory string
}

type BoxService struct {
	ProcessEnvValues   func(map[string]interface{}, string) map[string]interface{}
	ValidateBoxes      func([]Box, string) error
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
