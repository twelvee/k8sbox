package model

import (
	"fmt"
	"os"

	"github.com/k8s-box/k8sbox/internal/k8sbox"
	"github.com/k8s-box/k8sbox/pkg/k8sbox/structs"
)

func RunEnvironment(tomlFile string) (structs.Environment, error) {
	environment, runDirectory := lookForEnvironmentStep(tomlFile)
	validateEnvironmentStep(environment)
	validateBoxesStep(environment, runDirectory)
	result := deployEnvironmentStep(environment)

	fmt.Println("Aight we're done here!")
	return result, nil
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
	err := k8sbox.GetEnvironmentService().ValidateEnvironment(environment)
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

func deployEnvironmentStep(environment structs.Environment) structs.Environment {
	fmt.Print("Deploying...")
	result, err := k8sbox.GetEnvironmentService().DeployEnvironment(environment)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "Reasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
	return result
}
