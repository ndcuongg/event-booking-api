package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ndcuongg/event-booking-api.git/middlewares"
)

func RegisterRouter(server *gin.Engine) {
	// Events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/cancel", cancelRegistration)

	// Users
	server.POST("/signup", signup)
	server.POST("/login", login)
}
