package run

func getExample() string {
	return `
	boxie run --file /examples/environments/example_environment.toml // Rolls out the environment based on the toml specification

	boxie run -f /examples/environments/example_environment.toml // Rolls out the environment based on the toml specification
	`
}
