package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/dependency-injection-demo/handler"
)

func NewRouter(todoHandler *handler.TodoHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/todos", todoHandler.Add)
	r.GET("/todos", todoHandler.GetAll)
	r.DELETE("/todos/:id", todoHandler.Delete)

	return r
}
