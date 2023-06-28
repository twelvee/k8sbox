// Package describe is an entry point and help tools for describe command
package describe

func getExample() string {
	return `
	boxie describe environment {EvnironmentID} -n default // will describe the state of the environment by reference to its ID

	boxie describe env {EvnironmentID} --namespace default // will describe the state of the environment by reference to its ID
	`
}
