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
	router.GET("/posts/popular", controllers.GetPostsByViews())
	router.POST("/posts/:postId/comments", controllers.CreateComment())
	router.GET("/posts/:postId/comments", controllers.GetComments())
	router.PUT("/posts/:postId/comments/:commentId", controllers.EditComment())
}
