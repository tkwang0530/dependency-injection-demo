package di

import (
	"github.com/tkwang0530/dependency-injection-demo/repository"
	"go.uber.org/dig"
)

func ProvideRepositories(container *dig.Container) {
	container.Provide(repository.NewTodoRepository)
}
