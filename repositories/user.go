package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

type User struct{}

func NewUserRepository(i do.Injector) (*User, error) {
	return &User{}, nil
}

func (r *User) Get(c *gin.Context) {

}
