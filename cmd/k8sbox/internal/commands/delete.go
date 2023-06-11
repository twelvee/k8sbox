// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

// NewDeleteCommand is delete command entry point
func NewDeleteCommand() *cobra.Command {
	var (
		command   *cobra.Command
		namespace string

		getExample = `
		k8sbox delete environment {EnvironmentID} -n test // will delete the environment by reference to its ID

		k8sbox delete env {EnvironmentID} --namespace=default // will delete the environment by reference to its ID
		`
	)
	command = &cobra.Command{
		Use:     "delete",
		Short:   "Delete a resource",
		Long:    "Remove the resource from your k8s cluster. Use requires specifying the type of the resource as the first argument, as well as one of the flags indicating that the resource belongs to it.",
		Example: getExample,
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleDeleteCommand(command.Context(), args[0], args[1], namespace)
			return nil
		},
	}
	command.Flags().StringVarP(&namespace, "namespace", "n", "default", "The ID of the resource to be deleted.")
	return command
}
