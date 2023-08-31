package main

import (
	"MicroGo/configs"
	"MicroGo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	configs.ConnectDB()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "active",
			"author": "https://github.com/Fillonit",
			"github": "https://github.com/Fillonit/MicroGo",
		})
	})

	routes.UserRoute(router)
	routes.PostRoute(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  "404",
			"message": "Page not found",
		})
	})

	router.Run("localhost:8080")
}
