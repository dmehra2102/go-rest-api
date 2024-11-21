package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.POST("/events", createPost)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", updateEvent)
}