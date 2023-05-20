package k8sbox

import (
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
)

func GetEnvironmentStruct() structs.Environment {
	return structs.Environment{}
}

func GetBoxStruct() structs.Box {
	return structs.Box{}
}

func GetApplicationStruct() structs.Application {
	return structs.Application{}
}
