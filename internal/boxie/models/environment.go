// Package model is used as an model entry point
package model

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/twelvee/boxie/internal/boxie"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"k8s.io/utils/strings/slices"
)

var s = spinner.New(spinner.CharSets[21], 100*time.Millisecond)

// RunEnvironmentAsync will run the environment async
func RunEnvironmentAsync(environment structs.Environment, c chan bool) {
	err := RunEnvironment(environment)
	if err != nil {
		c <- false
	}
	c <- true
	close(c)
}

// RunEnvironment will prepare and deploy environment to your k8s cluster
func RunEnvironment(environment structs.Environment) error {
	start := time.Now()
	s.Start()
	if len(strings.TrimSpace(environment.TempDeployDirectory)) == 0 {
		err := createTempDirectoryStep(&environment)
		if err != nil {
			return err
		}
	}
	defer os.RemoveAll(environment.TempDeployDirectory)
	if len(strings.TrimSpace(environment.LoadBoxesFrom)) != 0 {
		err := loadBoxesStep(&environment)
		if err != nil {
			return err
		}
	}
	expandEnvironmentVariablesStep(&environment)
	expandBoxVariablesStep(&environment)
	err := validateEnvironmentStep(&environment)
	if err != nil {
		return err
	}
	err = validateBoxesStep(&environment)
	if err != nil {
		return err
	}
	kubeconfig := os.Getenv("KUBECONFIG")
	if len(strings.TrimSpace(environment.ClusterName)) != 0 {
		shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
		cluster, err := shelf.GetCluster(structs.GetClusterRequest{Name: environment.ClusterName})
		if err != nil {
			return err
		}
		kb, err := boxie.GetEnvironmentService().CreateTempKubeconfig(environment, cluster)
		if err != nil {
			return err
		}
		kubeconfig = kb
		// Delete kubeconfig only if it custom from cluster in shelf
		defer os.Remove(kb)
	}
	err = boxie.GetEnvironmentService().PrepareToWorkWithNamespace(environment.Namespace, kubeconfig)
	if err != nil {
		return err
	}
	err = removeLegacyEnvironmentStep(&environment)
	if err != nil {
		return err
	}
	err = deployEnvironmentStep(&environment)
	if err != nil {
		return err
	}
	s.Stop()
	fmt.Println("Alright, we're done here!")
	fmt.Printf("It took %.2fs.\r\n", time.Since(start).Seconds())
	return nil
}

// DeleteEnvironmentByID will delete saved environment by environmentID
func DeleteEnvironmentByID(namespace string, environmentName string) error {
	environment, err := boxie.GetStorageService().GetEnvironment(namespace, environmentName)
	if err != nil {
		return err
	}
	return deleteEnvironment(environment)
}

// DeleteEnvironmentByTomlFile will delete saved environment by initial toml file
func DeleteEnvironmentByTomlFile(namespace string, tomlFile string) error {
	environment, err := lookForEnvironmentStep(tomlFile)
	if err != nil {
		return err
	}
	return deleteEnvironment(&environment)
}

func deleteEnvironment(environment *structs.Environment) error {
	start := time.Now()
	expandEnvironmentVariablesStep(environment)
	expandBoxVariablesStep(environment)
	err := deleteEnvironmentStep(environment)
	if err != nil {
		return err
	}

	fmt.Println("Alright, we're done here!")
	fmt.Printf("It took %.2fs.\r\n", time.Since(start).Seconds())
	return nil
}

func createTempDirectoryStep(environment *structs.Environment) error {
	s.Suffix = " Creating work directory..."
	err := boxie.GetEnvironmentService().CreateTempDir(environment)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return fmt.Errorf("Failed: \n\r%s\n\r", err)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
	return nil
}

func lookForEnvironmentStep(tomlFile string) (structs.Environment, error) {
	s.Suffix = " Looking for the environment..."
	environment, err := boxie.GetTomlFormatter().GetEnvironmentFromToml(tomlFile)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return structs.Environment{}, fmt.Errorf("Failed: \n\r%s\n\r", err)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
	return environment, nil
}

func loadBoxesStep(environment *structs.Environment) error {
	s.Suffix = " Download boxes..."
	u, err := url.Parse(environment.LoadBoxesFrom)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return fmt.Errorf("Failed: \n\r%s\n\r", err)
	}
	if !slices.Contains(structs.GetAvailableDownloadSchemes(), u.Scheme) {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return fmt.Errorf("Failed: \n\r%s\n\r", err)
	}
	newEnvironment, err := boxie.GetTomlFormatter().GetEnvironmentViaHTTP(environment.LoadBoxesFrom, environment.LoadBoxesHeaders)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return fmt.Errorf("Failed: \n\r%s\n\r", err)
	}

	environment.Boxes = newEnvironment.Boxes
	return nil
}

func expandEnvironmentVariablesStep(environment *structs.Environment) {
	s.Suffix = " Expanding environment variables..."
	boxie.GetEnvironmentService().ExpandVariables(environment)
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
}

func expandBoxVariablesStep(environment *structs.Environment) {
	s.Suffix = " Expanding box variables..."
	environment.Boxes = boxie.GetBoxService().ExpandBoxVariables(environment.Boxes)
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
}

func validateEnvironmentStep(environment *structs.Environment) error {
	s.Suffix = " Validating the environment..."
	err := boxie.GetEnvironmentService().ValidateEnvironment(environment)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return fmt.Errorf("Failed: \n\r%s\n\r", err)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
	return nil
}

func validateBoxesStep(environment *structs.Environment) error {
	s.Suffix = " Validating boxes..."
	err := boxie.GetBoxService().ValidateBoxes(environment.Boxes)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return fmt.Errorf("Failed: \n\r%s\n\r", err)
	}

	for i := range environment.Boxes {
		err = boxie.GetBoxService().FillEmptyFields(*environment, &environment.Boxes[i])
		if err != nil {
			s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
			s.Stop()
			return fmt.Errorf("Failed: \n\r%s\n\r", err)
		}
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
	return nil
}

func removeLegacyEnvironmentStep(environment *structs.Environment) error {
	saved, err := boxie.GetStorageService().IsEnvironmentSaved(*environment)
	if err != nil {
		return err
	}
	if !saved {
		return nil
	}
	s.Suffix = " Deleting previous environment..."
	err = boxie.GetEnvironmentService().DeleteEnvironment(environment)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return fmt.Errorf("Failed: \n\r%s\n\r", err)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
	return nil
}

func deployEnvironmentStep(environment *structs.Environment) error {
	s.Suffix = " Deploying..."
	err := boxie.GetEnvironmentService().DeployEnvironment(environment)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return fmt.Errorf("Failed: \n\r%s\n\r", err)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
	return nil
}

func deleteEnvironmentStep(environment *structs.Environment) error {
	s.Suffix = " Deleting..."
	err := boxie.GetEnvironmentService().DeleteEnvironment(environment)
	if err != nil {
		s.Suffix = strings.Join([]string{s.Suffix, "FAIL"}, " ")
		s.Stop()
		return fmt.Errorf("Failed: \n\r%s\n\r", err)
	}
	s.Suffix = strings.Join([]string{s.Suffix, "OK"}, " ")
	return nil
}
