package handlers

import (
	"002_sample-go-app/usecases"
	"002_sample-go-app/usecases/interfaces"
	"github.com/gin-gonic/gin"
)

type User struct {
	Handler interfaces.User
}

func NewUserH(user *usecases.User) (*User, error) {
	return &User{
		Handler: user,
	}, nil
}

func (h *User) H(r *gin.RouterGroup) {
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//TODO
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "pong",
			"id":      id,
		})
	})
	//TODO
	r.POST("/user", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "pong",
			"id":      id,
		})
	})
	//TODO
	r.PATCH("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "pong",
			"id":      id,
		})
	})
}
