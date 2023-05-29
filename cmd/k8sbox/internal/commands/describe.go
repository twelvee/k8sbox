package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

func NewDescribeCommand() *cobra.Command {
	var (
		command *cobra.Command
	)
	command = &cobra.Command{
		Use:   "describe",
		Short: "Get detailed information about saved environment",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleDescribeCommand(command.Context(), args[0], args[1])
			return nil
		},
	}
	return command
}
