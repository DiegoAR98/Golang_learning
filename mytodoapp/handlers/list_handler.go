package handlers

import (
	"github.com/gin-gonic/gin"
	"example.com/mytodoapp/models"
	"net/http"
)

func ListTasks(c *gin.Context) {
	c.JSON(http.StatusOK, models.Tasks)
}
