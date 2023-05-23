package structs

type Application struct {
	Name          string
	Chart         string
	TempDirectory string
}

type ApplicationService struct {
	ValidateApplications func([]Application) []string
	ExpandApplications   func([]Application) []Application
}
