package k8sbox

import (
	"github.com/twelvee/k8sbox/internal/k8sbox/formatters"
	"github.com/twelvee/k8sbox/internal/k8sbox/services"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
)

func GetEnvironmentService() structs.EnvironmentService {
	return services.NewEnvironmentService()
}

func GetBoxService() structs.BoxService {
	return services.NewBoxService()
}

func GetApplicationService() structs.ApplicationService {
	return services.NewApplicationService()
}

func GetTomlFormatter() formatters.TomlFormatter {
	return formatters.NewTomlFormatter()
}
