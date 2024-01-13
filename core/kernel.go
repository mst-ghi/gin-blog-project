package core

import (
	"blog/database/models"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) models.User {
	return c.MustGet(gin.AuthUserKey).(models.User)
}

func Token(c *gin.Context) string {
	return c.MustGet("token").(string)
}
