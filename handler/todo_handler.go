package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/dependency-injection-demo/service"
)

type TodoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) Add(c *gin.Context) {
	var task struct {
		Task string `json:"task"`
	}
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := h.service.Add(task.Task)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	todos := h.service.GetAll()
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	h.service.Delete(id)
	c.Status(http.StatusNoContent)
}
