package articles

import (
	"blog/core"

	"github.com/gin-gonic/gin"
)

type ArticlesController struct {
	root    *core.Controller
	service ArticlesServiceInterface
}

func NewArticlesController() *ArticlesController {
	return &ArticlesController{
		root:    core.GetController(),
		service: NewArticlesService(),
	}
}

// @tags    Articles
// @router  /api/v1/articles [get]
// @summary get list of articles
// @accept  json
// @produce json
// @success 200 {object} core.Response[ArticlesResponseType]
func (self *ArticlesController) FindAll(c *gin.Context) {
	articles := self.service.FindAll()
	self.root.Success(c, ArticlesResponse(articles))
}

// @tags    Articles
// @router  /api/v1/articles/{id} [get]
// @summary get article by id
// @accept  json
// @produce json
// @success 200 {object} core.Response[ArticleResponseType]
// @param   id path int true "Article ID"
func (self *ArticlesController) FindOne(c *gin.Context) {
	article, err := self.service.FindOne(c.Param("id"))

	if err != nil {
		self.root.NotFoundError(c, err)
		return
	}

	self.root.Success(c, ArticleResponse(article))
}
