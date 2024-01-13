package app

import (
	"blog/app/articles"
	"blog/app/auth"
	"blog/app/users"
	"blog/core"

	"github.com/gin-gonic/gin"
)

// @tags App
// @router	/api [get]
// @summary	app route, get heathy status
func HomeRoute(c *gin.Context) {
	ctrl := core.GetController()
	ctrl.Success(c, map[string]string{
		"heathy": "I'm OK...",
	})
}

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", HomeRoute)

	v1Group := router.Group("/v1")
	{
		auth.RegisterRoutes(v1Group)
		users.RegisterRoutes(v1Group)
		articles.RegisterRoutes(v1Group)
	}
}
