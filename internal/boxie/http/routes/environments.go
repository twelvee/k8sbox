// Package routes contains every available REST API route
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie/http/controllers"
	"github.com/twelvee/boxie/internal/boxie/http/middlewares"
)

func initEnvironmentsRoutes(rg *gin.RouterGroup) {
	// Auth area
	rg.GET("/environments", middlewares.HasToken(), controllers.GetEnvironments)
	rg.GET("/environments/:name", middlewares.HasToken(), controllers.GetEnvironment)
	rg.POST("/environments", middlewares.HasToken(), controllers.CreateEnvironment)
	rg.DELETE("/environments/:name", middlewares.HasToken(), controllers.DeleteEnvironment)
}
