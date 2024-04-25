package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/commands", controllers.GetCommands)
	r.POST("/command", controllers.AddCommand)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	return r
}
