package usecases

import (
	"002_sample-go-app/repositories/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

type User struct {
	repository *interfaces.User
}

func NewUserUseCase(i do.Injector) (*User, error) {
	return &User{
		repository: do.MustInvoke[*interfaces.User](i),
	}, nil
}

func (u *User) Get(c *gin.Context) {

}
