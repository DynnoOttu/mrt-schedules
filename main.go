package main

import (
	"os"

	"github.com/DynnoOttu/mrt-schedules/modules/stations"
	"github.com/gin-gonic/gin"
)

func main() {
	initiateRouter()
}

func initiateRouter() {

	var (
		router = gin.Default()
		api    = router.Group("/api")
	)

	stations.Initiate(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
