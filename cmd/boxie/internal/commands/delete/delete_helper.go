// Package delete is an entry point and help tools for delete commands
package delete

func getExample() string {
	return `
	boxie delete environment {EnvironmentID} -n test // will delete the environment by reference to its ID

	boxie delete env {EnvironmentID} --namespace=default // will delete the environment by reference to its ID
	`
}

func getShelfExample() string {
	return `
	boxie shelf delete box first-box

	boxie shelf delete box my-box
	`
}
