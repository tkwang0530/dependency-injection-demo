package todo

import (
	"github.com/tkwang0530/dependency-injection-demo/models"
	"github.com/tkwang0530/dependency-injection-demo/repository/todo"
)

type impl struct {
	repo todo.Repository
}

func New(repo todo.Repository) Service {
	return &impl{repo: repo}
}

func (im *impl) Add(task string) int {
	return im.repo.Add(task)
}

func (im *impl) GetAll() []models.Todo {
	return im.repo.GetAll()
}

func (im *impl) Delete(id int) {
	im.repo.Delete(id)
}
