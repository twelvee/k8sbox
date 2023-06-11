// Package k8sbox is exporting all services and formatters
package k8sbox

import (
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
)

// GetEnvironmentStruct will return an environment struct
func GetEnvironmentStruct() structs.Environment {
	return structs.Environment{}
}

// GetBoxStruct will return a box struct
func GetBoxStruct() structs.Box {
	return structs.Box{}
}

// GetApplicationStruct will return an application struct
func GetApplicationStruct() structs.Application {
	return structs.Application{}
}

// GetStorageStruct will return an storage struct
func GetStorageStruct() structs.Storage {
	return structs.Storage{}
}
