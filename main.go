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

	routes.UserRoute(router)
	routes.PostRoute(router)

	router.Run("localhost:8080")
}
