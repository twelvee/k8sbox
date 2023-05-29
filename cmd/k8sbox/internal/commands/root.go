// Package commands is an entrypoint for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
)

var (
	root *cobra.Command
)

// Root command entrypoint
func NewRootCommand() *cobra.Command {
	root = &cobra.Command{
		Use:   "k8sbox [command] [flags]",
		Short: "k8sbox - cli tool that helps you to manage your k8s environments",
	}

	root.AddCommand(NewRunCommand())
	root.AddCommand(NewGetCommand())
	root.AddCommand(NewDeleteCommand())
	root.AddCommand(NewDescribeCommand())

	return root
}
