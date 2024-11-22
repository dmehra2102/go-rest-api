package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// User Routes
	server.POST("/login", login)
	server.POST("/signup", signup)

	// Events Routes
	server.GET("/events", getEvents)
	server.POST("/events", createPost)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
