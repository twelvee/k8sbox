package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

func NewRunCommand() *cobra.Command {
	var (
		command  *cobra.Command
		tomlFile string
	)
	command = &cobra.Command{
		Use:   "run",
		Short: "Run templated environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleRunCommand(command.Context(), tomlFile)
			return nil
		},
	}
	command.Flags().StringVarP(&tomlFile, "file", "f", "", "Environment template file (toml)")
	command.MarkFlagRequired("file")
	return command
}
