package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

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
			handlers.HandleGetCommand(args[0], flags, command.Context())
			return nil
		},
	}
	return command
}
