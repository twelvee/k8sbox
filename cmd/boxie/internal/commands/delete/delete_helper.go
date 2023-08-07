// Package delete is an entry point and help tools for delete commands
package delete

func getExample() string {
	return `
	boxie delete environment {name} -n test // will delete the environment by reference to its name

	boxie delete env {name} --namespace=default // will delete the environment by reference to its name
	`
}

func getShelfExample() string {
	return `
	boxie shelf delete box first-box

	boxie shelf delete box my-box
	`
}
