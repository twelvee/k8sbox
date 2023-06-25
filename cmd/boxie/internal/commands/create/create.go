// Package create is an entry point and help tools for create commands
package create

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/internal/boxie/handlers/get"
)

// NewCreateCommand is create command entry point
func NewCreateCommand() *cobra.Command {
	var (
		command   *cobra.Command
		namespace string
	)
	command = &cobra.Command{
		Use:     "create",
		Short:   "Create and save the resource",
		Long:    "Create and save the resource that you are going to use when creating environments.",
		Example: getExample(),
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			get.HandleGetCommand(command.Context(), args[0], namespace)
			return nil
		},
	}
	command.Flags().StringVarP(&namespace, "namespace", "n", "", "The namespace of the resource to be listed.")
	command.MarkFlagRequired("namespace")
	return command
}
