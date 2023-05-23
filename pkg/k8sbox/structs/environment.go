package structs

type Environment struct {
	Name          string
	Id            string
	Namespace     string
	Boxes         []Box
	TempDirectory string
	Variables     string
}

type EnvironmentService struct {
	DeployEnvironment         func(*Environment, bool) error
	DeleteEnvironment         func(*Environment) error
	CreateTempDeployDirectory func(*Environment, string, bool) (string, error)
	ValidateEnvironment       func(*Environment, string) error
	ExpandVariables           func(*Environment)
}
