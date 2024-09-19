package interfaces

import (
	"002_sample-go-app/models"
	"github.com/gin-gonic/gin"
)

type User interface {
	Get(c *gin.Context, id int) (*models.User, error)
}
