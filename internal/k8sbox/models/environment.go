// Package model is used as an model entrypoint
package model

import (
	"errors"
	"fmt"
	"os"

	"github.com/twelvee/k8sbox/internal/k8sbox"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
)

// RunEnvironment prepare and deploy environment to your k8s cluster
func RunEnvironment(tomlFile string) error {
	environment := lookForEnvironmentStep(tomlFile)
	expandEnvironmentVariablesStep(&environment)
	expandBoxVariablesStep(&environment)
	validateEnvironmentStep(&environment)
	validateBoxesStep(&environment)
	isSaved := checkIfEnvironmentIsSavedStep(environment)
	createTempDeployDirectoryStep(&environment, isSaved)
	if isSaved {
		checkIfEnvironmentHasSameBoxesStep(&environment)
	}
	deployEnvironmentStep(&environment, isSaved)

	fmt.Println("Aight we're done here!")
	return nil
}

// DeleteEnvironment prepare and uninstall environment from your k8s cluster
func DeleteEnvironment(tomlFile string) error {
	environment := lookForEnvironmentStep(tomlFile)
	expandEnvironmentVariablesStep(&environment)
	expandBoxVariablesStep(&environment)
	isSaved := checkIfEnvironmentIsSavedStep(environment)
	if !isSaved {
		return errors.New("Saved environment not found")
	}
	checkIfEnvironmentHasSameBoxesStep(&environment)
	deleteEnvironmentStep(&environment)

	fmt.Println("Aight we're done here!")
	return nil
}

func lookForEnvironmentStep(tomlFile string) structs.Environment {
	fmt.Print("Looking for environment...")
	environment, err := k8sbox.GetTomlFormatter().GetEnvironmentFromToml(tomlFile)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
	return environment
}

func expandEnvironmentVariablesStep(environment *structs.Environment) {
	fmt.Print("Expanding environment variables...")
	k8sbox.GetEnvironmentService().ExpandVariables(environment)
	fmt.Println(" OK")
}

func expandBoxVariablesStep(environment *structs.Environment) {
	fmt.Print("Expanding box variables...")
	environment.Boxes = k8sbox.GetBoxService().ExpandBoxVariables(environment.Boxes)
	fmt.Println(" OK")
}

func checkIfEnvironmentIsSavedStep(environment structs.Environment) bool {
	fmt.Print("Matching with already saved environments...")
	saved, err := utils.IsEnvironmentSaved(environment.Id)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	if saved {
		fmt.Println(" OK - SAVED")
		return true
	}
	fmt.Println(" OK - NEW")
	return false
}

func checkIfEnvironmentHasSameBoxesStep(environment *structs.Environment) {
	fmt.Print("Matching boxes on founded environment...")
	savedEnvironment, err := utils.GetEnvironment(environment.Id)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
	if len(savedEnvironment.Boxes) > 0 {
		fmt.Printf("Found %d legacy boxes. Removing...", len(savedEnvironment.Boxes))

		for _, savedBox := range savedEnvironment.Boxes {
			_, err := k8sbox.GetBoxService().UninstallBox(&savedBox, *environment)
			if err != nil {
				fmt.Println(" FAIL :(")
				fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
				os.Exit(1)
			}
		}
		fmt.Println(" OK")
	}
}

func validateEnvironmentStep(environment *structs.Environment) {
	fmt.Print("Validating environment...")
	err := k8sbox.GetEnvironmentService().ValidateEnvironment(environment)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}

func validateBoxesStep(environment *structs.Environment) {
	fmt.Print("Validating boxes...")
	err := k8sbox.GetBoxService().ValidateBoxes(environment.Boxes)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}

	for i, _ := range environment.Boxes {
		err = k8sbox.GetBoxService().FillEmptyFields(&environment.Boxes[i], environment.Namespace)
		if err != nil {
			fmt.Println(" FAIL :(")
			fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
			os.Exit(1)
		}
	}
	fmt.Println(" OK")
}

func createTempDeployDirectoryStep(environment *structs.Environment, isSaved bool) {
	fmt.Print("Moving files to a temporary directory...")
	var err error
	environment.TempDirectory, err = k8sbox.GetEnvironmentService().CreateTempDeployDirectory(environment, isSaved)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "\n\rReasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}

func deployEnvironmentStep(environment *structs.Environment, isSaved bool) {
	fmt.Print("Deploying...")
	err := k8sbox.GetEnvironmentService().DeployEnvironment(environment, isSaved)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "Reasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}

func deleteEnvironmentStep(environment *structs.Environment) {
	fmt.Print("Deleting environment...")
	err := k8sbox.GetEnvironmentService().DeleteEnvironment(environment)
	if err != nil {
		fmt.Println(" FAIL :(")
		fmt.Fprintf(os.Stderr, "Reasons: \n\r%s\n\r", err)
		os.Exit(1)
	}
	fmt.Println(" OK")
}
