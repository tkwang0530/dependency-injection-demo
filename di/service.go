package di

import (
	"github.com/tkwang0530/dependency-injection-demo/service/todo"
	"go.uber.org/dig"
)

func ProvideServices(container *dig.Container) {
	container.Provide(todo.New)
}
