package structs

type Environment struct {
	Name      string
	Tag       string
	Namespace string
	Services  []string
}

type EnvironmentService struct {
	DeployEnvironment func(Environment) (Environment, error)
}
