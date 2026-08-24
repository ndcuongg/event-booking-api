package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ndcuongg/event-booking-api.git/db"
	"github.com/ndcuongg/event-booking-api.git/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRouter(server)
	server.Run(":8080")
}
