package listeners

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TorHttpRouter() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Standard HTTP Server is running on port 42069",
			},
		)
	})

	return e
}
