package articles

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewArticlesController()

	guestGroup := router.Group("/articles")
	{
		guestGroup.GET("", ctrl.FindAll)
		guestGroup.GET("/:id", ctrl.FindOne)
	}
}
