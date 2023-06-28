// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/cmd/boxie/internal/commands/delete"
	"github.com/twelvee/boxie/cmd/boxie/internal/commands/describe"
	"github.com/twelvee/boxie/cmd/boxie/internal/commands/get"
	"github.com/twelvee/boxie/cmd/boxie/internal/commands/run"
)

var (
	root *cobra.Command
)

// NewRootCommand is root command entry point
func NewRootCommand() *cobra.Command {
	root = &cobra.Command{
		Use:   "boxie [command] [flags]",
		Short: "boxie - A tool that allows you to roll out your environments into your k8s cluster",
		Long:  "boxie - A tool that allows you to roll out your environments into your k8s cluster using templated specifications, monitor the activity of these services, as well as easily clean up the cluster of unused resources that you rolled out earlier. ",
	}

	root.AddCommand(run.NewRunCommand())
	root.AddCommand(get.NewGetCommand())
	root.AddCommand(delete.NewDeleteCommand())
	root.AddCommand(describe.NewDescribeCommand())
	root.AddCommand(NewShelfCommand())

	return root
}
