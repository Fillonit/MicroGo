package routes

import (
	"MicroGo/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {
	router.POST("/product", controllers.CreateProduct())
	router.GET("/product/:productId", controllers.GetProduct())
	router.PUT("/product/:productId", controllers.UpdateProduct())
	router.DELETE("/product/:productId", controllers.DeleteProduct())
	router.GET("/products", controllers.GetAllProducts())
}
