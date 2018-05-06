package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllTodos allows to get all todos
func GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not Implemented yet",
	})
}

// GetOneTodo allows to get One todo
func GetOneTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not Implemented yet",
	})
}

// AddOneTodo allows to add one todo
func AddOneTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not Implemented yet",
	})
}

// EditOneTodo allows to edit one todo
func EditOneTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not Implemented yet",
	})
}

// DeleteOneTodo allows to delete one todo
func DeleteOneTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Not Implemented yet",
	})
}
