// Package structs contain every k8sbox public structs
package structs

// Application is your box application in a struct
type Application struct {
	Name          string `toml:"name"`
	Chart         string `toml:"chart"`
	TempDirectory string `toml:"-"`
}

// ApplicationService is a public ApplicationService
type ApplicationService struct {
	ValidateApplications func([]Application) []string
	ExpandApplications   func([]Application) []Application
}

// GetApplicationAliases return a slice of application model name aliases
func GetApplicationAliases() []string {
	return []string{"application", "applications", "apps"}
}
