// Package structs contain every boxie public structs
package structs

// Environment is your environment in a struct
type Environment struct {
	Name                    string                   `toml:"name"`
	Namespace               string                   `toml:"namespace"`
	Boxes                   []Box                    `toml:"boxes"`
	Variables               string                   `toml:"variables"`
	VariablesMap            map[string]string        `toml:"variables_map"`
	LoadBoxesFrom           string                   `toml:"load_boxes_from"`
	LoadBoxesHeaders        map[string]Header        `toml:"load_boxes_headers"`
	TempDeployDirectory     string                   `toml:"-"`
	ClusterName             string                   `toml:"-"`
	EnvironmentApplications []EnvironmentApplication `toml:"-"`
	UserID                  int32                    `toml:"-"`
	CreatedAt               string                   `toml:"-"`
	Status                  int32                    `toml:"-"`
}

// EnvironmentService is a public EnvironmentService
type EnvironmentService struct {
	DeployEnvironment          func(*Environment) error
	DeleteEnvironment          func(*Environment) error
	ValidateEnvironment        func(*Environment) error
	ExpandVariables            func(*Environment)
	PrepareToWorkWithNamespace func(namespace string, kubeconfig string) error
	CreateTempDir              func(*Environment) error
	CreateTempKubeconfig       func(Environment, Cluster) (string, error)
	FillWithRuntimeData        func(*Environment) error
}

// GetEnvironmentAliases return a slice of environment model name aliases
func GetEnvironmentAliases() []string {
	return []string{"environment", "environments", "env"}
}

// GetAvailableDownloadSchemes will return all downloadable schemas as slice
func GetAvailableDownloadSchemes() []string {
	return []string{"http", "https"}
}

const (
	ENVIRONMENT_STATUS_PENDING    int32 = 0
	ENVIRONMENT_STATUS_INSTALLING int32 = 1
	ENVIRONMENT_STATUS_RUNNING    int32 = 2
	ENVIRONMENT_STATUS_FAILED     int32 = 3
	ENVIRONMENT_STATUS_STOPPED    int32 = 4
)

// DeleteEnvironmentRequest is environment delete rest api request
type DeleteEnvironmentRequest struct {
	Name string
}

// UpdateEnvironmentStatusRequest is environment update status rest api request
type UpdateEnvironmentStatusRequest struct {
	Name   string
	Status int32
}

// GetEnvironmentRequest is get environment by name rest api request
type GetEnvironmentRequest struct {
	Name string
}
