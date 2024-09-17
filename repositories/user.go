package repositories

import (
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

func (r *User) Get(c *gin.Context) {

}
