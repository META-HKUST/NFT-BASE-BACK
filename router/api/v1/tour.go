package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tutorials struct {
	Articles []Article
}

type Article struct {
	ArticleId int `json:"article_id"`
}

// @Description  tour
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/tour"
// @Router       /tour [GET]
func GetAllTutorials(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  tutorials
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/tr/articles"
// @Router       /tr/articles [GET]
func GetAllArticles(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  view articles by title
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/tr/articles/XXXXXX"
// @Router      /tr/articles/:articles-id [POST]
func GetArticleByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  event-banner
// @Tags         tour
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/event-banner"
// @Router       /event-banner [GET]
func Name19(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}
