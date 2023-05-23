package structs

type Environment struct {
	Name          string `toml:"name"`
	Id            string `toml:"id"`
	Namespace     string `toml:"namespace"`
	Boxes         []Box  `toml:"boxes"`
	TempDirectory string `toml:"-"`
	Variables     string `toml:"variables"`
}

type EnvironmentService struct {
	DeployEnvironment         func(*Environment, bool) error
	DeleteEnvironment         func(*Environment) error
	CreateTempDeployDirectory func(*Environment, bool) (string, error)
	ValidateEnvironment       func(*Environment) error
	ExpandVariables           func(*Environment)
}
