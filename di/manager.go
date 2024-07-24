package di

import (
	"github.com/tkwang0530/dependency-injection-demo/router"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	ProvideConfig(container)
	ProvideDatabase(container)
	ProvideRepositories(container)
	ProvideServices(container)
	ProvideHandlers(container)

	container.Provide(router.NewRouter)

	return container
}
