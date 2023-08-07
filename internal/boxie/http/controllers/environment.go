// Package controllers contains every REST API route logic
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie"
	model "github.com/twelvee/boxie/internal/boxie/models"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"net/http"
	"os"
	"strings"
)

// GetEnvironments will return a serialized environments list struct as json
func GetEnvironments(c *gin.Context) {
	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	environments, err := shelf.GetEnvironments()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, environment := range environments {
		cluster, err := shelf.GetCluster(structs.GetClusterRequest{Name: environment.ClusterName})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		kubeconfig, err := boxie.GetEnvironmentService().CreateTempKubeconfig(environment, cluster)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer os.Remove(kubeconfig)

		err = boxie.GetEnvironmentService().PrepareToWorkWithNamespace(environment.Namespace, kubeconfig)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = boxie.GetEnvironmentService().FillWithRuntimeData(&environments[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"environments": environments})
}

// DeleteEnvironment will delete environment by its name
func DeleteEnvironment(c *gin.Context) {
	var input structs.DeleteEnvironmentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Name = c.Param("name")
	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))

	environment, err := shelf.GetEnvironment(structs.GetEnvironmentRequest{Name: input.Name})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	cluster, err := shelf.GetCluster(structs.GetClusterRequest{Name: environment.ClusterName})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	kubeconfig, err := boxie.GetEnvironmentService().CreateTempKubeconfig(environment, cluster)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer os.Remove(kubeconfig)

	err = boxie.GetEnvironmentService().PrepareToWorkWithNamespace(environment.Namespace, kubeconfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var boxes []structs.Box
	for i, _ := range environment.EnvironmentApplications {
		inside := false
		for j, b := range boxes {
			if b.Name == environment.EnvironmentApplications[i].BoxName {
				// Box already inside slice
				boxes[j].HelmRender[environment.EnvironmentApplications[i].Name] = environment.EnvironmentApplications[i].Chart
				inside = true
				continue
			}
		}
		if !inside {
			rend := make(map[string]string)
			var box structs.Box
			box.Name = environment.EnvironmentApplications[i].BoxName
			// TODO: replace with actual type
			box.Type = structs.Helm()
			box.Namespace = environment.Namespace
			rend[environment.EnvironmentApplications[i].Name] = environment.EnvironmentApplications[i].Chart
			box.HelmRender = rend
			boxes = append(boxes, box)
		}
	}

	environment.Boxes = boxes

	err = boxie.GetEnvironmentService().DeleteEnvironment(&environment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = shelf.DeleteEnvironment(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// GetEnvironment will return serialized environment
func GetEnvironment(c *gin.Context) {
	var input structs.GetEnvironmentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Name = c.Param("name")

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	environment, err := shelf.GetEnvironment(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"environment": environment})
}

// CreateEnvironment will create and run environment
func CreateEnvironment(c *gin.Context) {
	var input structs.Environment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(strings.TrimSpace(input.TempDeployDirectory)) == 0 {
		err := boxie.GetEnvironmentService().CreateTempDir(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	err := boxie.GetEnvironmentService().ValidateEnvironment(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, _ := range input.Boxes {
		err = boxie.GetBoxService().FillEmptyFields(input, &input.Boxes[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	user, err := shelf.GetUser(c.GetHeader("x-auth-token"))
	if err != nil {
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
	}

	err = shelf.PutEnvironment(input, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"created": true})

	runEnvironment(input)
}

func runEnvironment(environment structs.Environment) {
	c := make(chan bool)
	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))

	go model.RunEnvironmentAsync(environment, c)
	fail := shelf.UpdateEnvironmentStatus(structs.UpdateEnvironmentStatusRequest{Status: structs.ENVIRONMENT_STATUS_INSTALLING, Name: environment.Name})
	if fail != nil {
		fmt.Println("Status update job was failed.", fail.Error())
	}
	select {
	case success := <-c:
		{
			if success {
				fail := shelf.UpdateEnvironmentStatus(structs.UpdateEnvironmentStatusRequest{Status: structs.ENVIRONMENT_STATUS_RUNNING, Name: environment.Name})
				if fail != nil {
					fmt.Println("Status update job was failed.", fail.Error())
				}
			} else {
				fail := shelf.UpdateEnvironmentStatus(structs.UpdateEnvironmentStatusRequest{Status: structs.ENVIRONMENT_STATUS_FAILED, Name: environment.Name})
				if fail != nil {
					fmt.Println("Status update job was failed.", fail.Error())
				}
			}
		}
	}
}
