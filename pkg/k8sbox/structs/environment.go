package structs

type Environment struct {
	Name      string
	Tag       string
	Namespace string
	Boxes     []Box
}

type EnvironmentService struct {
	DeployEnvironment func(Environment) (Environment, error)
}
