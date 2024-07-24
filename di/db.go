package di

import (
	"github.com/tkwang0530/dependency-injection-demo/db"
	"go.uber.org/dig"
)

func ProvideDatabase(container *dig.Container) {
	container.Provide(db.NewInMemoryDB)
}
