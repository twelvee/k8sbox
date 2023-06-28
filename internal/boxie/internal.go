// Package boxie is exporting all services and formatters
package boxie

import (
	"github.com/twelvee/boxie/internal/boxie/formatters"
	"github.com/twelvee/boxie/internal/boxie/services"
	"github.com/twelvee/boxie/internal/boxie/shelf"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

// GetEnvironmentService will put and return a new EnvironmentService
func GetEnvironmentService() structs.EnvironmentService {
	return services.NewEnvironmentService()
}

// GetBoxService will put and return a new BoxService
func GetBoxService() structs.BoxService {
	return services.NewBoxService()
}

// GetApplicationService will put and return a new ApplicationService
func GetApplicationService() structs.ApplicationService {
	return services.NewApplicationService()
}

// GetStorageService will put and return a new StorageService
func GetStorageService() structs.StorageService {
	return services.NewStorageService()
}

// GetTomlFormatter will put and return a new TomlFormatter
func GetTomlFormatter() formatters.TomlFormatter {
	return formatters.NewTomlFormatter()
}

// GetJsonFormatter will put and return a new JsonFormatter
func GetJsonFormatter() formatters.JsonFormatter {
	return formatters.NewJsonFormatter()
}

// GetShelf will put and return a new Shelf instance
func GetShelf(connectionType string, dsn string) shelf.Shelf {
	return shelf.NewShelf(connectionType, dsn)
}
