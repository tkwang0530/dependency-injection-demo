package di

import (
	"github.com/tkwang0530/dependency-injection-demo/config"
	"go.uber.org/dig"
)

func ProvideConfig(container *dig.Container) {
	container.Provide(config.NewConfig)
}
