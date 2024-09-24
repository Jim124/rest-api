package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	// server.POST("/events", middlewares.Authenticate, createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)
	authorized := server.Group("/")
	authorized.Use(middlewares.Authenticate)
	authorized.POST("/events", createEvent)
	authorized.PUT("/events/:id", updateEvent)
	authorized.DELETE("/events/:id", deleteEvent)
	authorized.POST("/events/:id/register", registerForEvent)
	authorized.DELETE("/events/:id/cancel", cancelRegistration)
	server.GET("/events/:id", getSingleEvent)
	server.POST("/signUp", signUp)
	server.POST("/login", login)
}
