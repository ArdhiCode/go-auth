package main

import (
	"log"

	"github.com/ArdhiCode/go-auth/internal/config"
	"github.com/ArdhiCode/go-auth/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	config.ConnectMongo()

	r := gin.Default()

	routes.RegisterRoutes(r)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
