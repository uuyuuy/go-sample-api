package main

import (
	"002_sample-go-app/handlers"
	"002_sample-go-app/repositories"
	"002_sample-go-app/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	// DB
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASS"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// DI
	userRepository, err := repositories.NewUserRepository(db)
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
