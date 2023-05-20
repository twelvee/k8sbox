package model

import (
	"fmt"
	"os"

	"github.com/twelvee/k8sbox/internal/k8sbox"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
)

func RunEnvironment(tomlFile string) error {
	environment, runDirectory := lookForEnvironmentStep(tomlFile)
	validateEnvironmentStep(environment)
	validateBoxesStep(environment, runDirectory)
	createTempDeployDirectoryStep(&environment, runDirectory)
	deployEnvironmentStep(&environment)

	fmt.Println("Aight we're done here!")
	return nil
}

func lookForEnvironmentStep(tomlFile string) (structs.Environment, string) {
	fmt.Print("Looking for environment...")
	environment, runDirectory, err := k8sbox.GetTomlFormatter().GetEnvironmentFromToml(tomlFile)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
	return environment, runDirectory
}

func validateEnvironmentStep(environment structs.Environment) {
	fmt.Print("Validating environment...")
	err := k8sbox.GetEnvironmentService().ValidateEnvironment(&environment)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}

func validateBoxesStep(environment structs.Environment, runDirectory string) {
	fmt.Print("Validating boxes...")
	err := k8sbox.GetBoxService().ValidateBoxes(environment.Boxes, runDirectory, k8sbox.GetApplicationService())
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}

func createTempDeployDirectoryStep(environment *structs.Environment, runDirectory string) {
	fmt.Print("Moving files to a temporary directory...")
	var err error
	environment.TempDirectory, err = k8sbox.GetEnvironmentService().CreateTempDeployDirectory(environment, runDirectory, utils.GetShortID(8))
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}

func deployEnvironmentStep(environment *structs.Environment) {
	fmt.Print("Deploying...")
	err := k8sbox.GetEnvironmentService().DeployEnvironment(environment)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "Reasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}
