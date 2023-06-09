// Package model is used as an model entry point
package model

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/twelvee/k8sbox/internal/k8sbox"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"k8s.io/utils/strings/slices"
)

var s = spinner.New(spinner.CharSets[21], 100*time.Millisecond)

// RunEnvironment will prepare and deploy environment to your k8s cluster
func RunEnvironment(tomlFile string) error {
	start := time.Now()
	s.Start()
	environment := lookForEnvironmentStep(tomlFile)
	if len(strings.TrimSpace(environment.LoadBoxesFrom)) != 0 {
		loadBoxesStep(&environment)
	}
	expandEnvironmentVariablesStep(&environment)
	expandBoxVariablesStep(&environment)
	validateEnvironmentStep(&environment)
	validateBoxesStep(&environment)
	err := k8sbox.GetEnvironmentService().PrepareToWorkWithNamespace(environment.Namespace)
	if err != nil {
		return err
	}
	removeLegacyEnvironmentStep(&environment)
	deployEnvironmentStep(&environment)
	s.Stop()
	fmt.Println("Alright, we're done here!")
	fmt.Printf("It took %.2fs.\r\n", time.Since(start).Seconds())
	return nil
}

// DeleteEnvironmentByID will delete saved environment by environmentID
func DeleteEnvironmentByID(namespace string, environmentID string) error {
	environment, err := k8sbox.GetStorageService().GetEnvironment(namespace, environmentID)
	if err != nil {
		return err
	}
	return deleteEnvironment(environment)
}

// DeleteEnvironmentByTomlFile will delete saved environment by initial toml file
func DeleteEnvironmentByTomlFile(namespace string, tomlFile string) error {
	environment := lookForEnvironmentStep(tomlFile)
	return deleteEnvironment(&environment)
}

func deleteEnvironment(environment *structs.Environment) error {
	start := time.Now()
	expandEnvironmentVariablesStep(environment)
	expandBoxVariablesStep(environment)
	deleteEnvironmentStep(environment)

	fmt.Println("Alright, we're done here!")
	fmt.Printf("It took %.2fs.\r\n", time.Since(start).Seconds())
	return nil
}

func lookForEnvironmentStep(tomlFile string) structs.Environment {
	s.Suffix = " Looking for the environment..."
	environment, err := k8sbox.GetTomlFormatter().GetEnvironmentFromToml(tomlFile)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		fmt.Fprintf(os.Stderr, "Failed: \n\r%s\n\r", err)
		os.Exit(1)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
	return environment
}

func loadBoxesStep(environment *structs.Environment) {
	s.Suffix = " Download boxes..."
	u, err := url.Parse(environment.LoadBoxesFrom)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		fmt.Fprintf(os.Stderr, "Failed: \n\r%s\n\r", err)
		os.Exit(1)
	}
	if !slices.Contains(structs.GetAvailableDownloadSchemes(), u.Scheme) {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		fmt.Fprintf(os.Stderr, "Failed. Available load_boxes_from scheme is %s\n\r", strings.Join(structs.GetAvailableDownloadSchemes(), ", "))
		os.Exit(1)
	}
	newEnvironment, err := k8sbox.GetTomlFormatter().GetEnvironmentViaHTTP(environment.LoadBoxesFrom, environment.LoadBoxesHeaders)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		fmt.Fprintf(os.Stderr, "Failed: \n\r%s\n\r", err)
		os.Exit(1)
	}

	environment.Boxes = newEnvironment.Boxes
}

func expandEnvironmentVariablesStep(environment *structs.Environment) {
	s.Suffix = " Expanding environment variables..."
	k8sbox.GetEnvironmentService().ExpandVariables(environment)
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
}

func expandBoxVariablesStep(environment *structs.Environment) {
	s.Suffix = " Expanding box variables..."
	environment.Boxes = k8sbox.GetBoxService().ExpandBoxVariables(environment.Boxes)
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
}

func validateEnvironmentStep(environment *structs.Environment) {
	s.Suffix = " Validating the environment..."
	err := k8sbox.GetEnvironmentService().ValidateEnvironment(environment)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		fmt.Fprintf(os.Stderr, "Failed: \n\r%s\n\r", err)
		os.Exit(1)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
}

func validateBoxesStep(environment *structs.Environment) {
	s.Suffix = " Validating boxes..."
	err := k8sbox.GetBoxService().ValidateBoxes(environment.Boxes)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		fmt.Fprintf(os.Stderr, "Failed: \n\r%s\n\r", err)
		os.Exit(1)
	}

	for i := range environment.Boxes {
		err = k8sbox.GetBoxService().FillEmptyFields(*environment, &environment.Boxes[i])
		if err != nil {
			s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
			s.Stop()
			fmt.Fprintf(os.Stderr, "Failed: \n\r%s\n\r", err)
			os.Exit(1)
		}
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
}

func removeLegacyEnvironmentStep(environment *structs.Environment) {
	saved, _ := k8sbox.GetStorageService().IsEnvironmentSaved(*environment)
	if !saved {
		return
	}
	s.Suffix = " Deleting previous environment..."
	err := k8sbox.GetEnvironmentService().DeleteEnvironment(environment)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		fmt.Fprintf(os.Stderr, "Failed: \n\r%s\n\r", err)
		os.Exit(1)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
}

func deployEnvironmentStep(environment *structs.Environment) {
	s.Suffix = " Deploying..."
	err := k8sbox.GetEnvironmentService().DeployEnvironment(environment)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		fmt.Fprintf(os.Stderr, "Failed: \n\r%s\n\r", err)
		os.Exit(1)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
}

func deleteEnvironmentStep(environment *structs.Environment) {
	s.Suffix = " Deleting..."
	err := k8sbox.GetEnvironmentService().DeleteEnvironment(environment)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		fmt.Fprintf(os.Stderr, "Failed: \n\r%s\n\r", err)
		os.Exit(1)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
}
