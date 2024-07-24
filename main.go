package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/dependency-injection-demo/di"
)

func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(router *gin.Engine) {
		if err := router.Run(":8080"); err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
