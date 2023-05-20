package structs

type Application struct {
	Name          string
	Tag           string
	Chart         string
	TempDirectory string
}

type ApplicationService struct {
	ValidateApplications func([]Application, string) []string
}
