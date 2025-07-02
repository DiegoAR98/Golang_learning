package routes

import (
	"github.com/gin-gonic/gin"
	"example.com/mytodoapp/handlers"
)

func CreateRoutes(router *gin.Engine) {
	router.POST("/tasks", handlers.CreateTask)
} 
