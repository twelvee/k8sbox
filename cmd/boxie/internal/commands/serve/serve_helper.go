// Package serve is an entry point and help tools for serve commands
package serve

func getExample() string {
	return `
	boxie serve --port=8888 --addr=0.0.0.0

	boxie serve
	`
}
