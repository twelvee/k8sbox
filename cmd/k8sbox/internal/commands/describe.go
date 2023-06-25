// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/k8sbox/internal/k8sbox/handlers"
)

// NewDescribeCommand is describe command entry point
func NewDescribeCommand() *cobra.Command {
	var (
		command   *cobra.Command
		namespace string

		getExample = `
		k8sbox describe environment {EvnironmentID} -n default // will describe the state of the environment by reference to its ID

		k8sbox describe env {EvnironmentID} --namespace default // will describe the state of the environment by reference to its ID
		`
	)
	command = &cobra.Command{
		Use:     "describe",
		Short:   "Describes the state of the selected resource",
		Long:    "Describes the state of the resource in the k8s cluster. Use requires the resource type as the first argument, and one of the flags indicating that the resource belongs.",
		Example: getExample,
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			handlers.HandleDescribeCommand(command.Context(), args[0], args[1], namespace)
			return nil
		},
	}
	command.Flags().StringVarP(&namespace, "namespace", "n", "default", "The ID of the resource to be described.")
	command.MarkFlagRequired("namespace")
	return command
}
