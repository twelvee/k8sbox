// Package create is an entry point and help tools for create commands
package create

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/internal/boxie/handlers/create"
)

// NewCreateCommand is create command entry point
func NewCreateCommand() *cobra.Command {
	var (
		command   *cobra.Command
		arguments create.CreateArguments
	)
	command = &cobra.Command{
		Use:     "create",
		Short:   "Create and save the resource",
		Long:    "Create and save the resource that you are going to use when creating environments.",
		Example: getExample(),
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			create.HandleCreateCommand(command.Context(), args[0], arguments)
			return nil
		},
	}
	command.Flags().StringVarP(&arguments.Namespace, "namespace", "ns", "", "The namespace of the resource to be created.")
	command.Flags().StringVarP(&arguments.BoxName, "name", "n", "", "The name of the resource to be created.")
	command.Flags().StringVarP(&arguments.BoxType, "type", "t", "", "The type of the resource to be created.")
	command.Flags().StringVarP(&arguments.ChartLocation, "chart", "c", "", "The chart location of the resource to be created.")
	command.Flags().StringVarP(&arguments.ValuesLocation, "values", "v", "", "The values location of the resource to be created.")
	command.Flags().StringVarP(&arguments.JsonSpec, "json", "j", "", "The json specification of the resource to be created.")
	command.Flags().StringVarP(&arguments.TomlFile, "file", "f", "", "The toml file location of the resource to be created.")
	return command
}
