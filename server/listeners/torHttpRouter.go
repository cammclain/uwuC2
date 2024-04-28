package listeners

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TorHttpRouter() http.Handler {
	tor_http_gin_client := gin.New()
	tor_http_gin_client.Use(gin.Recovery())
	tor_http_gin_client.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Standard HTTP Server is running on port 42069",
			},
		)
	})

	return tor_http_gin_client
}
