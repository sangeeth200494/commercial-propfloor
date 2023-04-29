package main

import (
	"commercial-propfloor/routes"
	"log"

	// "os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8099"

	router := gin.New()
	router.Use(gin.Logger())
	// routes.PublicRoutes(router)
	routes.PrivateRoutes(router)

	log.Fatal(router.Run(":" + port))

}
