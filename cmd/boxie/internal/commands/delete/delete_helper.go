package delete

func getExample() string {
	return `
	boxie delete environment {EnvironmentID} -n test // will delete the environment by reference to its ID

	boxie delete env {EnvironmentID} --namespace=default // will delete the environment by reference to its ID
	`
}
