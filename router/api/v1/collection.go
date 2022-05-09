package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description  all collections
// @Tags         collection
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/collections"
// @Router       /collections [GET]
func Name10_1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  single collection
// @Tags         collection
// @param 		 collection-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/collections/hahahah"
// @Router       /collections/{collection-id} [GET]
func Name10_2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  all items in collection
// @Tags         collection
// @param 		 collection-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/collections/hahahah/items"
// @Router       /collections/{collection-id}/items [GET]
func Name10_3(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}
