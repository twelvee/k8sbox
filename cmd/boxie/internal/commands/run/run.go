// Package run is an entry point and help tools for run commands
package run

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/internal/boxie/handlers"
)

// NewRunCommand is run command entry point
func NewRunCommand() *cobra.Command {
	var (
		command  *cobra.Command
		tomlFile string
	)
	command = &cobra.Command{
		Use:     "run",
		Short:   "Run the environment",
		Long:    "Run the environment with the toml specification.",
		Example: getExample(),
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleRunCommand(command.Context(), tomlFile)
			return nil
		},
	}
	command.Flags().StringVarP(&tomlFile, "file", "f", "", "Path toml file specifying the environment to be created.")
	command.MarkFlagRequired("file")
	return command
}
