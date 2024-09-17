package main

import (
	"002_sample-go-app/handlers"
	"002_sample-go-app/repositories"
	"002_sample-go-app/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	// DI
	userRepository, err := repositories.NewUserRepository()
	if err != nil {
		return
	}
	userUseCase, err := usecases.NewUserUseCase(userRepository)
	if err != nil {
		return
	}
	userHandler, err := handlers.NewUserH(userUseCase)
	if err != nil {
		return
	}

	// Router
	r := gin.Default()
	group := r.Group("/v1")

	userHandler.H(group)
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
