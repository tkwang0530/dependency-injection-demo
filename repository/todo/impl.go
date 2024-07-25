package todo

import (
	"github.com/tkwang0530/dependency-injection-demo/db"
	"github.com/tkwang0530/dependency-injection-demo/models"
)

type impl struct {
	db *db.InMemoryDB
}

func New(db *db.InMemoryDB) Repository {
	return &impl{db: db}
}

func (im *impl) Add(task string) int {
	return im.db.Add(task)
}

func (im *impl) GetAll() []models.Todo {
	todos := im.db.GetAll()
	result := make([]models.Todo, 0, len(todos))
	for id, task := range todos {
		result = append(result, models.Todo{ID: id, Task: task})
	}
	return result
}

func (im *impl) Delete(id int) {
	im.db.Delete(id)
}
