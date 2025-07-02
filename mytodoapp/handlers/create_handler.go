package handlers

import (
	"github.com/gin-gonic/gin"
	"example.com/mytodoapp/models"
	"net/http"
)

func CreateTask (c *gin.Context) {
	var task models.Task

	// Bind the JSON input to the Task struct
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign an ID (in a real application, this would be handled by a database)
	task.ID = len(models.Tasks) + 1
	task.Done = false

	// Add the new task to the list
	models.Tasks = append(models.Tasks, task)

	c.JSON(http.StatusCreated, task)
}
