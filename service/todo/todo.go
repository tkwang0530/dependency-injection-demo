package todo

import "github.com/tkwang0530/dependency-injection-demo/models"

type Service interface {
	Add(task string) int
	GetAll() []models.Todo
	Delete(id int)
}
