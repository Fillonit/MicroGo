package routes

import (
	"MicroGo/controllers"
	"MicroGo/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
	router.PUT("/user/:userId", middlewares.AuthMiddleware(), controllers.EditAUser())
	router.DELETE("/user/:userId", middlewares.AuthMiddleware(), controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
}
