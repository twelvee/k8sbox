package structs

type Box struct {
	Tag           string
	Type          string
	Applications  []Application
	Chart         string
	Values        string
	Namespace     string
	Name          string
	TempDirectory string
}

type BoxService struct {
	ValidateBoxes func([]Box, string, ApplicationService) error
}

func Helm() string {
	return "helm"
}

func Plain() string {
	return "plain"
}
