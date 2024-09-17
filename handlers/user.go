package handlers

import (
	"002_sample-go-app/usecases"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

type User struct {
	Handler *usecases.User
}

func NewUserH(i do.Injector) (*User, error) {
	return &User{
		Handler: do.MustInvoke[*usecases.User](i),
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
