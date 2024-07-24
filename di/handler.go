package di

import (
	"github.com/tkwang0530/dependency-injection-demo/handler"
	"go.uber.org/dig"
)

func ProvideHandlers(container *dig.Container) {
	container.Provide(handler.NewTodoHandler)
}
