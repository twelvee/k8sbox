// Package commands is an entrypoint for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

// NewDeleteCommand is delete command entrypoint
func NewDeleteCommand() *cobra.Command {
	var (
		command  *cobra.Command
		tomlFile string
		id       string
	)
	command = &cobra.Command{
		Use:   "delete",
		Short: "Uninstall model",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleDeleteCommand(command.Context(), args[0], tomlFile, id)
			return nil
		},
	}
	command.Flags().StringVarP(&id, "id", "i", "", "Model ID")
	command.Flags().StringVarP(&tomlFile, "file", "f", "", "Environment template file (toml)")
	return command
}
