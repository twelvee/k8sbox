// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

// NewRunCommand is run command entry point
func NewRunCommand() *cobra.Command {
	var (
		command  *cobra.Command
		tomlFile string
	)
	command = &cobra.Command{
		Use:   "run",
		Short: "Run the environment",
		Long:  "Run the environment with the toml specification.",
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleRunCommand(command.Context(), tomlFile)
			return nil
		},
	}
	command.Flags().StringVarP(&tomlFile, "file", "f", "", "Path toml file specifying the environment to be created.")
	command.MarkFlagRequired("file")
	return command
}
