// Package structs contain every boxie public structs
package structs

// Storage is your environment storage as a struct
type Storage struct {
	Type StorageType `toml:"type"`
}

// StorageType is an enum that has all available storage types
type StorageType string

const (
	TYPE_FILESYSTEM StorageType = "filesystem"
	TYPE_VOLUME     StorageType = "volume"
)

// StorageService is a public StorageService
type StorageService struct {
	EnsureStorageAvailable func(string) error
	SaveEnvironment        func(Environment) error
	DeleteEnvironment      func(Environment) error
	DeleteBox              func(Environment, Box) error
	GetEnvironments        func(string) ([]Environment, error)
	GetEnvironment         func(namespace string, id string) (*Environment, error)
	IsEnvironmentSaved     func(Environment) (bool, error)
}
