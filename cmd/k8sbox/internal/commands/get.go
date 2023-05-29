// Package commands is an entrypoint for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

// NewGetCommand is get command entrypoint
func NewGetCommand() *cobra.Command {
	var (
		command *cobra.Command
		flags   []string
	)
	command = &cobra.Command{
		Use:   "get",
		Short: "Get saved environments",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleGetCommand(command.Context(), args[0], flags)
			return nil
		},
	}
	return command
}
