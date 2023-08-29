package routes

import (
	"MicroGo/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoute(router *gin.Engine) {
	router.POST("/post", controllers.CreatePost())
	router.GET("/post/:postId", controllers.GetAPost())
	router.PUT("/post/:postId", controllers.EditAPost())
	router.DELETE("/post/:postId", controllers.DeleteAPost())
	router.GET("/posts", controllers.GetAllPosts())
}
