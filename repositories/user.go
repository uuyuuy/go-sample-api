package repositories

import (
	"002_sample-go-app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (*User, error) {
	return &User{
		db: db,
	}, nil
}

func (r *User) Get(c *gin.Context, id int) (*models.User, error) {
	user := &models.User{}
	res := r.db.First(user, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
