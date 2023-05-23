package structs

type Environment struct {
	Name          string
	Id            string
	Namespace     string
	Boxes         []Box
	TempDirectory string
}

type EnvironmentService struct {
	DeployEnvironment         func(*Environment, bool) error
	DeleteEnvironment         func(*Environment) error
	CreateTempDeployDirectory func(*Environment, string, bool) (string, error)
	ValidateEnvironment       func(*Environment) error
}
