// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
)

var (
	root *cobra.Command
)

// NewRootCommand is root command entry point
func NewRootCommand() *cobra.Command {
	root = &cobra.Command{
		Use:   "k8sbox [command] [flags]",
		Short: "k8sbox - A tool that allows you to roll out your environments into your k8s cluster",
		Long:  "k8sbox - A tool that allows you to roll out your environments into your k8s cluster using templated specifications, monitor the activity of these services, as well as easily clean up the cluster of unused resources that you rolled out earlier. ",
	}

	root.AddCommand(NewRunCommand())
	root.AddCommand(NewGetCommand())
	root.AddCommand(NewDeleteCommand())
	root.AddCommand(NewDescribeCommand())

	return root
}
