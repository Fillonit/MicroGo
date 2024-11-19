package main

import (
	"MicroGo/configs"
	"MicroGo/routes"

	"github.com/gin-gonic/gin"
	nrgin "github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("MicroGo"),
		newrelic.ConfigLicense("eu01xx94bbd13030341078e873dba86dFFFFNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(nrgin.Middleware(app))

	router.SetTrustedProxies([]string{"127.0.0.1"})
	configs.EnvMongoURI()
	configs.EnvSecretKey()
	configs.ConnectDB()
	gin.SetMode(gin.ReleaseMode)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "active",
			"author": "https://github.com/Fillonit",
			"github": "https://github.com/Fillonit/MicroGo",
		})
	})

	routes.UserRoute(router)
	routes.PostRoute(router)
	routes.ProductRoute(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  "404",
			"message": "Page not found",
		})
	})

	router.Run("0.0.0.0:8080")
}
