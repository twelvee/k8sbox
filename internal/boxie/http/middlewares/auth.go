// Package middlewares contains every REST API middlewares
package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/twelvee/boxie/internal/boxie"
	"net/http"
	"os"
	"strings"
)

func HasToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		token := c.GetHeader("x-auth-token")
		if len(strings.TrimSpace(token)) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "x-auth-token header is required.")
			return
		}

		shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
		_, err := shelf.GetUser(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		c.Next() // request
		// after request
	}
}
