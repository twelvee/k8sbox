// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

// NewGetCommand is get command entry point
func NewGetCommand() *cobra.Command {
	var (
		command   *cobra.Command
		namespace string

		getExample = `
		k8sbox get environments // get a list of saved environments

		k8sbox get environment // get a list of saved environments

		k8sbox get env // get a list of saved environments
		`
	)
	command = &cobra.Command{
		Use:     "get",
		Short:   "Get a list of saved resources",
		Long:    "Get a list of earlier created resources. Use requires the resource type as the first argument.",
		Example: getExample,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleGetCommand(command.Context(), args[0], namespace)
			return nil
		},
	}
	command.Flags().StringVarP(&namespace, "namespace", "n", "", "The namespace of the resource to be listed.")
	command.MarkFlagRequired("namespace")
	return command
}
