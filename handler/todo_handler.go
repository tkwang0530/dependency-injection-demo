package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/dependency-injection-demo/service/todo"
)

type TodoHandler struct {
	todoService todo.Service
}

func NewTodoHandler(todoService todo.Service) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (h *TodoHandler) Add(c *gin.Context) {
	var task struct {
		Task string `json:"task"`
	}
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := h.todoService.Add(task.Task)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	todos := h.todoService.GetAll()
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	h.todoService.Delete(id)
	c.Status(http.StatusNoContent)
}
