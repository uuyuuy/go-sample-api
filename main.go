package main

import (
	"002_sample-go-app/handlers"
	"002_sample-go-app/repositories"
	"002_sample-go-app/usecases"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
	"log"
)

func main() {
	// DI Container
	injector := do.New()
	do.Provide(injector, repositories.NewUserRepository)
	do.Provide(injector, usecases.NewUserUseCase)
	do.Provide(injector, handlers.NewUserH)

	// Router
	r := gin.Default()
	group := r.Group("/v1")
	h, err := do.Invoke[handlers.User](injector)
	if err != nil {
		log.Fatal(err.Error())
	}
	h.H(group)
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
