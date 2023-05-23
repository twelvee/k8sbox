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
	ValidateBoxes   func([]Box, string) error
	FillEmptyFields func(*Box, string) error
	UninstallBox    func(*Box, string) (*release.UninstallReleaseResponse, error)
}

func Helm() string {
	return "helm"
}

func Plain() string {
	return "plain"
}
