package main

import (
	"gitlab.com/drxos/rest-api/todos/urls"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	urls.Todo(r)
	r.Run(":8000")
}
