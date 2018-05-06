package urls

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/drxos/rest-api/todos/handlers"
)

// Todo routes
func Todo(r *gin.Engine) *gin.Engine {
	r.GET("/todos", handlers.GetAllTodos)
	r.GET("/todos/:id", handlers.GetOneTodo)
	r.POST("/todos/", handlers.AddOneTodo)
	r.PUT("/todos/:id", handlers.EditOneTodo)
	r.DELETE("/todos/:id", handlers.DeleteOneTodo)
	return r
}
