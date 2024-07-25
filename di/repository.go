package di

import (
	"github.com/tkwang0530/dependency-injection-demo/repository/todo"
	"go.uber.org/dig"
)

func ProvideRepositories(container *dig.Container) {
	container.Provide(todo.New)
}
