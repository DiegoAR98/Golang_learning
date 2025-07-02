package handlers

import (
	"github.com/gin-gonic/gin"
	"example.com/mytodoapp/models"
	"net/http"
	"strconv"
)

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, task := range models.Tasks {
		if strconv.Itoa(task.ID) == id {
			// Remove the task from the slice
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
