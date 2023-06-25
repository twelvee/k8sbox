// Package handlers is used to process Cobra commands
package handlers

import (
	"context"

	"github.com/twelvee/boxie/internal/boxie"
)

// KuberExecutable prepare all data to use k8s cluster with provided flags
func KuberExecutable(context context.Context, namespace string) {
	err := boxie.GetEnvironmentService().PrepareToWorkWithNamespace(namespace)
	if err != nil {
		panic(err)
	}
}
