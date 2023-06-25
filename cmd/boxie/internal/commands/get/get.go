// Package get is an entry point and help tools for get commands
package get

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/internal/boxie/handlers"
)

// NewGetCommand is get command entry point
func NewGetCommand() *cobra.Command {
	var (
		command   *cobra.Command
		namespace string
	)
	command = &cobra.Command{
		Use:     "get",
		Short:   "Get a list of saved resources",
		Long:    "Get a list of earlier created resources. Use requires the resource type as the first argument.",
		Example: getExample(),
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
