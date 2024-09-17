package interfaces

import "github.com/gin-gonic/gin"

type User interface {
	Get(c *gin.Context)
}
