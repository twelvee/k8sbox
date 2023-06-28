// Package structs contain every boxie public structs
package structs

// Environment is your environment in a struct
type Environment struct {
	Name                string            `toml:"name"`
	ID                  string            `toml:"id"`
	Namespace           string            `toml:"namespace"`
	Boxes               []Box             `toml:"boxes"`
	Variables           string            `toml:"variables"`
	VariablesMap        map[string]string `toml:"variables_map"`
	LoadBoxesFrom       string            `toml:"load_boxes_from"`
	LoadBoxesHeaders    map[string]Header `toml:"load_boxes_headers"`
	TempDeployDirectory string            `toml:"-"`
}

// EnvironmentService is a public EnvironmentService
type EnvironmentService struct {
	DeployEnvironment          func(*Environment) error
	DeleteEnvironment          func(*Environment) error
	ValidateEnvironment        func(*Environment) error
	ExpandVariables            func(*Environment)
	PrepareToWorkWithNamespace func(namespace string) error
	CreateTempDir              func(*Environment) error
}

// GetEnvironmentAliases return a slice of environment model name aliases
func GetEnvironmentAliases() []string {
	return []string{"environment", "environments", "env"}
}

func GetAvailableDownloadSchemes() []string {
	return []string{"http", "https"}
}
