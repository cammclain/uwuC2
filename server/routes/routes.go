package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/commands", controllers.GetCommands)
	r.POST("/command", controllers.AddCommand)

	return r
}
