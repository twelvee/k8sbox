// Package get is an entry point and help tools for get commands
package get

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/internal/boxie/handlers/get"
)

// NewShelfGetCommand is get command entry point
func NewShelfGetCommand() *cobra.Command {
	var (
		command *cobra.Command
	)
	command = &cobra.Command{
		Use:     "get",
		Short:   "Get a list of saved resources",
		Long:    "Get a list of earlier created resources. Use requires the resource type as the first argument.",
		Example: getShelfExample(),
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			get.HandleShelfGetCommand(command.Context(), args[0])
			return nil
		},
	}
	return command
}
