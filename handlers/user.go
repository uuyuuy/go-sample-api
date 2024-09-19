package handlers

import (
	"002_sample-go-app/usecases"
	"002_sample-go-app/usecases/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	useCase interfaces.User
}

func NewUserH(user *usecases.User) (*User, error) {
	return &User{
		useCase: user,
	}, nil
}

func (h *User) H(r *gin.RouterGroup) {
	//TODO
	r.GET("/users", func(c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
			})
		}
		h.useCase.Get(c, id)
	})
	r.GET("/user/:id", func(c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		c.Header("Content-Type", "application/json; charset=utf-8")
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
			})
		}
		user, err := h.useCase.Get(c, id)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
			})
		}
		c.JSON(200, gin.H{
			"message": "success",
			"user":    user,
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
