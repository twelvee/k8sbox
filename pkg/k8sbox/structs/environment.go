package structs

type Environment struct {
	Name      string
	Tag       string
	Namespace string
	Boxes     []Box
}

type EnvironmentService struct {
	DeployEnvironment         func(Environment, string) (Environment, error)
	CreateTempDeployDirectory func(Environment, string, string) (string, error)
	ValidateEnvironment       func(Environment) error
}
