package main

import (
	"github.com/brunoleitem/rest-api-go/db"
	"github.com/brunoleitem/rest-api-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":3333")
}
