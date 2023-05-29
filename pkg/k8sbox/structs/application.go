package structs

type Application struct {
	Name          string `toml:"name"`
	Chart         string `toml:"chart"`
	TempDirectory string `toml:"-"`
}

type ApplicationService struct {
	ValidateApplications func([]Application) []string
	ExpandApplications   func([]Application) []Application
}
