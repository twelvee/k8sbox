package commands

import (
	"github.com/k8s-box/k8sbox/internal/k8sbox/handlers"
	"github.com/spf13/cobra"
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
			handlers.HandleRunCommand(tomlFile, command.Context())
			return nil
		},
	}
	command.Flags().StringVarP(&tomlFile, "file", "f", "", "Environment template file (toml)")
	command.MarkFlagRequired("file")
	return command
}
