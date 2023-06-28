// Package describe is an entry point and help tools for describe command
package describe

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/internal/boxie/handlers/describe"
)

// NewDescribeCommand is describe command entry point
func NewDescribeCommand() *cobra.Command {
	var (
		command   *cobra.Command
		namespace string
	)
	command = &cobra.Command{
		Use:     "describe",
		Short:   "Describes the state of the selected resource",
		Long:    "Describes the state of the resource in the k8s cluster. Use requires the resource type as the first argument, and one of the flags indicating that the resource belongs.",
		Example: getExample(),
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			describe.HandleDescribeCommand(command.Context(), args[0], args[1], namespace)
			return nil
		},
	}
	command.Flags().StringVarP(&namespace, "namespace", "n", "default", "The ID of the resource to be described.")
	command.MarkFlagRequired("namespace")
	return command
}
