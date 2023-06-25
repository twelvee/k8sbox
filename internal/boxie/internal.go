// Package boxie is exporting all services and formatters
package boxie

import (
	"github.com/twelvee/boxie/internal/boxie/formatters"
	"github.com/twelvee/boxie/internal/boxie/services"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

// GetEnvironmentService will create and return a new EnvironmentService
func GetEnvironmentService() structs.EnvironmentService {
	return services.NewEnvironmentService()
}

// GetBoxService will create and return a new BoxService
func GetBoxService() structs.BoxService {
	return services.NewBoxService()
}

// GetApplicationService will create and return a new ApplicationService
func GetApplicationService() structs.ApplicationService {
	return services.NewApplicationService()
}

// GetStorageService will create and return a new StorageService
func GetStorageService() structs.StorageService {
	return services.NewStorageService()
}

// GetTomlFormatter will create and return a new TomlFormatter
func GetTomlFormatter() formatters.TomlFormatter {
	return formatters.NewTomlFormatter()
}
