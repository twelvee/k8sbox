// Package run is used to process run command
package run

// RunArguments is run arguments request as a struct
type RunArguments struct {
	TomlFile      string
	Boxes         []string
	Name          string
	Namespace     string
	VariablesPath string
	Variables     map[string]string
	ID            string
}
