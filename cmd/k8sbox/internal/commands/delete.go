package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

func NewDeleteCommand() *cobra.Command {
	var (
		command  *cobra.Command
		tomlFile string
	)
	command = &cobra.Command{
		Use:   "delete",
		Short: "Uninstall environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleDeleteCommand(tomlFile, command.Context())
			return nil
		},
	}
	command.Flags().StringVarP(&tomlFile, "file", "f", "", "Environment template file (toml)")
	command.MarkFlagRequired("file")
	return command
}
