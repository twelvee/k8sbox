// Package routes contains every available REST API route
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie/http/controllers"
)

func initSetupRoutes(rg *gin.RouterGroup) {
	rg.GET("/setup_required", controllers.GetSetupRequired)
	rg.POST("/user/create_first", controllers.Setup)
}
