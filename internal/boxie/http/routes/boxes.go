// Package routes contains every available REST API route
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie/http/controllers"
	"github.com/twelvee/boxie/internal/boxie/http/middlewares"
)

func initBoxesRoutes(rg *gin.RouterGroup) {
	// Auth area
	rg.GET("/boxes", middlewares.HasToken(), controllers.GetBoxes)
	rg.GET("/boxes/:name", middlewares.HasToken(), controllers.GetBox)
	rg.PUT("/boxes/:name", middlewares.HasToken(), controllers.UpdateBox)
	rg.POST("/boxes", middlewares.HasToken(), controllers.CreateBox)
	rg.DELETE("/boxes/:name", middlewares.HasToken(), controllers.DeleteBox)
}
