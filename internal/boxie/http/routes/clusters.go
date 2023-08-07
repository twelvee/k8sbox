// Package routes contains every available REST API route
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie/http/controllers"
	"github.com/twelvee/boxie/internal/boxie/http/middlewares"
)

func initClustersRoutes(rg *gin.RouterGroup) {
	// Auth area
	rg.GET("/clusters", middlewares.HasToken(), controllers.GetClusters)
	rg.GET("/clusters/:name", middlewares.HasToken(), controllers.GetCluster)
	rg.PUT("/clusters/:name", middlewares.HasToken(), controllers.UpdateCluster)
	rg.POST("/clusters", middlewares.HasToken(), controllers.CreateCluster)
	rg.POST("/clusters/test", middlewares.HasToken(), controllers.TestClusterConnection)
	rg.DELETE("/clusters/:name", middlewares.HasToken(), controllers.DeleteCluster)
}
