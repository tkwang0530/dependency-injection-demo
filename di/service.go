package di

import (
	"github.com/tkwang0530/dependency-injection-demo/service"
	"go.uber.org/dig"
)

func ProvideServices(container *dig.Container) {
	container.Provide(service.NewTodoService)
}
