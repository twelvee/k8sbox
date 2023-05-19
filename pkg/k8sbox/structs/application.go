package structs

type Application struct {
	Name  string
	Tag   string
	Chart string
}

type ApplicationService struct {
	ValidateApplications func([]Application, string) []string
}
