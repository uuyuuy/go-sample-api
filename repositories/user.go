package repositories

import (
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUserRepository() (*User, error) {
	return &User{}, nil
}

func (r *User) Get(c *gin.Context) {

}
