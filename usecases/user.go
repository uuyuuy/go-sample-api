package usecases

import (
	"002_sample-go-app/models"
	"002_sample-go-app/repositories"
	"002_sample-go-app/repositories/interfaces"
	"github.com/gin-gonic/gin"
)

type User struct {
	repository interfaces.User
}

func NewUserUseCase(repository *repositories.User) (*User, error) {
	return &User{
		repository: repository,
	}, nil
}

func (u *User) Get(c *gin.Context, id int) (*models.User, error) {
	user, err := u.repository.Get(c, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
