package structs

type Environment struct {
	Name          string
	Tag           string
	Namespace     string
	Boxes         []Box
	TempDirectory string
}

type EnvironmentService struct {
	DeployEnvironment         func(*Environment) error
	CreateTempDeployDirectory func(*Environment, string, string) (string, error)
	ValidateEnvironment       func(*Environment) error
}
