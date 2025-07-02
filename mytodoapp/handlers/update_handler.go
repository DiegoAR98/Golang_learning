package handlers

import (
	"github.com/gin-gonic/gin"
	"example.com/mytodoapp/models"
	"net/http"
	"strconv"
)

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task

	// Bind the JSON input to the Task struct
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range models.Tasks {
		if strconv.Itoa(task.ID) == id {
			// Update the task details
			models.Tasks[i].Title = updatedTask.Title
			models.Tasks[i].Description = updatedTask.Description
			models.Tasks[i].Done = updatedTask.Done

			c.JSON(http.StatusOK, models.Tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}