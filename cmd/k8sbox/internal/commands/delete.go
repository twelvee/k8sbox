// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

// NewDeleteCommand is delete command entry point
func NewDeleteCommand() *cobra.Command {
	var (
		command  *cobra.Command
		tomlFile string
		id       string

		getExample = `
		k8sbox delete environment -i {EnvironmentID} // will delete the environment by reference to its ID

		k8sbox delete env -i {EnvironmentID} // will delete the environment by reference to its ID

		k8sbox delete environment --file ./examples/environments/example_environment.toml // will delete the environment by reference to its toml specification

		k8sbox delete env --file ./examples/environments/example_environment.toml // will delete the environment by reference to its toml specification
		`
	)
	command = &cobra.Command{
		Use:     "delete",
		Short:   "Delete a resource",
		Long:    "Remove the resource from your k8s cluster. Use requires specifying the type of the resource as the first argument, as well as one of the flags indicating that the resource belongs to it.",
		Example: getExample,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleDeleteCommand(command.Context(), args[0], tomlFile, id)
			return nil
		},
	}
	command.Flags().StringVarP(&id, "id", "i", "", "The ID of the resource to be deleted.")
	command.Flags().StringVarP(&tomlFile, "file", "f", "", "Path toml file specifying the environment to be deleted.")
	return command
}
