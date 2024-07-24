package service

import (
	"github.com/tkwang0530/dependency-injection-demo/models"
	"github.com/tkwang0530/dependency-injection-demo/repository"
)

type TodoService interface {
	Add(task string) int
	GetAll() []models.Todo
	Delete(id int)
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) Add(task string) int {
	return s.repo.Add(task)
}

func (s *todoService) GetAll() []models.Todo {
	return s.repo.GetAll()
}

func (s *todoService) Delete(id int) {
	s.repo.Delete(id)
}
