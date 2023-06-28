// Package run is an entry point and help tools for run commands
package run

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/internal/boxie/handlers/run"
)

// NewRunCommand is run command entry point
func NewRunCommand() *cobra.Command {
	var (
		command *cobra.Command
		runArgs run.RunArguments
	)
	command = &cobra.Command{
		Use:     "run",
		Short:   "Run the environment",
		Long:    "Run the environment with the toml specification.",
		Example: getExample(),
		RunE: func(cmd *cobra.Command, args []string) error {
			run.HandleRunCommand(command.Context(), runArgs)
			return nil
		},
	}
	command.Flags().StringVarP(&runArgs.TomlFile, "file", "f", "", "Path toml file specifying the environment to be created.")
	command.Flags().StringArrayVarP(&runArgs.Boxes, "box", "b", nil, "Environment boxes.")
	command.Flags().StringToStringVarP(&runArgs.Variables, "env", "e", nil, "Environment variables.")
	command.Flags().StringVarP(&runArgs.VariablesPath, "env-path", "", "", "Environment variables file path.")
	command.Flags().StringVarP(&runArgs.Name, "name", "n", "", "Environment name. Should be unique.")
	command.Flags().StringVarP(&runArgs.Namespace, "namespace", "", "default", "Environment namespace.")
	command.Flags().StringVarP(&runArgs.ID, "id", "i", "", "Environment ID.")

	return command
}
