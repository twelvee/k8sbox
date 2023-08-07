// Package commands is an entry point for every single cobra command available
package commands

import (
	"github.com/spf13/cobra"
	"github.com/twelvee/boxie/cmd/boxie/internal/commands/delete"
	"github.com/twelvee/boxie/cmd/boxie/internal/commands/get"
	"github.com/twelvee/boxie/cmd/boxie/internal/commands/put"
)

var (
	shelf *cobra.Command
)

// NewShelfCommand is shelf command entry point
func NewShelfCommand() *cobra.Command {
	shelf = &cobra.Command{
		Use:   "shelf [command] [flags]",
		Short: "shelf - A database that stores your boxes",
		Long:  "shelf - A database that securely stores your boxes and with which you can put new environments through the web version of boxie.",
	}

	shelf.AddCommand(put.NewPutCommand())
	shelf.AddCommand(delete.NewShelfDeleteCommand())
	shelf.AddCommand(get.NewShelfGetCommand())

	return shelf
}
