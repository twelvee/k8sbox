// Package serve is an entry point and help tools for serve commands
package serve

import (
	"github.com/spf13/cobra"
	serve "github.com/twelvee/boxie/internal/boxie/handlers/serve"
)

// NewServeCommand is serve command entry point
func NewServeCommand() *cobra.Command {
	var (
		command *cobra.Command
		host    string
		port    string
	)
	command = &cobra.Command{
		Use:     "serve",
		Short:   "Serve boxie rest api and UI app2",
		Long:    "Launches the Boxie in server mode, as well as running the UI interface for convenient use.",
		Example: getExample(),
		RunE: func(cmd *cobra.Command, args []string) error {
			serve.HandleServeCommand(command.Context(), host, port)
			return nil
		},
	}
	command.Flags().StringVarP(&port, "port", "p", "8888", "Server port.")
	command.Flags().StringVarP(&host, "addr", "a", "0.0.0.0", "Server host.")

	return command
}
