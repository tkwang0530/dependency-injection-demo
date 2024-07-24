package repository

import (
	"github.com/tkwang0530/dependency-injection-demo/db"
	"github.com/tkwang0530/dependency-injection-demo/models"
)

type TodoRepository interface {
	Add(task string) int
	GetAll() []models.Todo
	Delete(id int)
}

type todoRepository struct {
	db *db.InMemoryDB
}

func NewTodoRepository(db *db.InMemoryDB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Add(task string) int {
	return r.db.Add(task)
}

func (r *todoRepository) GetAll() []models.Todo {
	todos := r.db.GetAll()
	result := make([]models.Todo, 0, len(todos))
	for id, task := range todos {
		result = append(result, models.Todo{ID: id, Task: task})
	}
	return result
}

func (r *todoRepository) Delete(id int) {
	r.db.Delete(id)
}
