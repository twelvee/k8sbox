// Package delete is an entry point and help tools for delete commands
package delete

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/internal/boxie/handlers/delete"
)

// NewShelfDeleteCommand is shelf delete command entry point
func NewShelfDeleteCommand() *cobra.Command {
	var (
		command *cobra.Command
	)
	command = &cobra.Command{
		Use:     "delete",
		Short:   "Delete a resource from shelf",
		Long:    "Remove the resource from your shelf storage. Use requires specifying the type of the resource as the first argument, as well as one of the flags indicating that the resource belongs to it.",
		Example: getShelfExample(),
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			delete.HandleShelfDeleteCommand(command.Context(), args[0], args[1])
			return nil
		},
	}
	return command
}
