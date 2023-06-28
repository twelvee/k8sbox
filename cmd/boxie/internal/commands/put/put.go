// Package put is an entry point and help tools for put commands
package put

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/internal/boxie/handlers/put"
)

// NewPutCommand is put command entry point
func NewPutCommand() *cobra.Command {
	var (
		command   *cobra.Command
		arguments put.PutArguments
		force     bool
	)
	command = &cobra.Command{
		Use:     "put",
		Short:   "Put resource into a shelf",
		Long:    "Put resource that you are going to use when creating environments into a shelf.",
		Example: getExample(),
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			put.HandlePutCommand(command.Context(), args[0], arguments, force)
			return nil
		},
	}
	command.Flags().StringVarP(&arguments.Namespace, "namespace", "", "", "The namespace of the resource to be created.")
	command.Flags().StringVarP(&arguments.BoxName, "name", "n", "", "The name of the resource to be created.")
	command.Flags().StringVarP(&arguments.BoxType, "type", "t", "", "The type of the resource to be created.")
	command.Flags().StringVarP(&arguments.ChartLocation, "chart", "c", "", "The chart location of the resource to be created.")
	command.Flags().StringVarP(&arguments.ValuesLocation, "values", "v", "", "The values location of the resource to be created.")
	command.Flags().StringVarP(&arguments.JsonSpec, "json", "j", "", "The json specification of the resource to be created.")
	command.Flags().StringVarP(&arguments.TomlFile, "file", "f", "", "The toml file location of the resource to be created.")
	command.Flags().BoolVarP(&force, "force", "", false, "Override box?")
	return command
}
