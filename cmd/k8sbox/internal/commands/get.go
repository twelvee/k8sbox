// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

// NewGetCommand is get command entry point
func NewGetCommand() *cobra.Command {
	var (
		command *cobra.Command
		flags   []string
	)
	command = &cobra.Command{
		Use:   "get",
		Short: "Get a list of saved resources",
		Long:  "Get a list of earlier created resources. Use requires the resource type as the first argument.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleGetCommand(command.Context(), args[0], flags)
			return nil
		},
	}
	return command
}
