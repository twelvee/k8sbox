// Package run is an entry point and help tools for run commands
package run

func getExample() string {
	return `
	boxie run --file /examples/environments/example_environment.toml // Rolls out the environment based on the toml specification

	boxie run -f /examples/environments/example_environment.toml // Rolls out the environment based on the toml specification
	`
}
