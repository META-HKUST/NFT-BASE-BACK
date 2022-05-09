package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description  get user
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna"
// @Router       /{id} [GET]
func Name1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  collected
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna/collected"
// @Router       /{id}/collected [GET]
func Name2(ctx *gin.Context) {
	//

	//

	//
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  favorites
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna/favorites"
// @Router       /{id}/favorites [GET]
func Name3(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  creation
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna/creation"
// @Router       /{id}/creation [GET]
func Name4(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  change profile
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 userName   query   string   false   "user Name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/change-profile"
// @Router       /{id}/change-profile [POST]
func Name5(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  create-collection
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 collection-Name   query   string   false   "collection Name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/create-collection"
// @Router       /{id}/create-collection [POST]
func Name9_1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  edit-collection
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 collection-Name   query   string   false   "collection Name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/edit-collection"
// @Router       /{id}/edit-collection [POST]
func Name9_2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  create-item
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 item-Name   query   string   false   "item Name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/create-item"
// @Router       /{id}/create-item [POST]
func Name15_1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  edit-item
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 item-Name   query   string   false   "item Name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/edit-item"
// @Router       /{id}/edit-item [POST]
func Name15_2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}
