package main

import (
	"github.com/gin-gonic/gin"
	"shestapalau.by/rest/db"
	"shestapalau.by/rest/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
