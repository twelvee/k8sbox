// Package controllers contains every REST API route logic
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"net/http"
	"os"
)

// GetClusters will return a serialized clusters list struct as json
func GetClusters(c *gin.Context) {
	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	clusters, err := shelf.GetClusters()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"clusters": clusters})
}

// DeleteCluster will delete cluster by its name
func DeleteCluster(c *gin.Context) {
	var input structs.DeleteClusterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Name = c.Param("name")

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	err := shelf.DeleteCluster(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

// GetCluster will return serialized cluster
func GetCluster(c *gin.Context) {
	var input structs.GetClusterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Name = c.Param("name")

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	cluster, err := shelf.GetCluster(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cluster": cluster})
}

// UpdateCluster will update cluster by its name
func UpdateCluster(c *gin.Context) {
	var input structs.Cluster
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.Name = c.Param("name")

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	isActive, err := boxie.GetClusterService().TestConnection(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.IsActive = isActive
	err = shelf.UpdateCluster(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated": true})
}

// TestClusterConnection will return connection status of given kubeconfig
func TestClusterConnection(c *gin.Context) {
	var input structs.Cluster
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAvailable, err := boxie.GetClusterService().TestConnection(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available": isAvailable})
}

// CreateCluster will create cluster
func CreateCluster(c *gin.Context) {
	var input structs.Cluster
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	err := shelf.PutCluster(input, false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"created": true})
}
