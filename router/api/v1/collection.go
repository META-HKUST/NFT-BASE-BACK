package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
type Collection struct {
	CollectionId          int       `json:"collection_id"`
	CollectionName        string    `json:"collection_name"`
	Items			  	  []Item
	CreateTime        	  time.Time `json:"create_time"`
	Owner				  string	`json:"owner"`
}

// @Description  get all collections under the current account
// @Tags         collection
// @param 		 collection-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/collections"
// @Router       /collections [GET]
func GetAllCollectionsByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  create new collection
// @Tags         collection
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST /api/v1/collections/create"
// @Router       /create [POST]
func CreateCollection(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}


// @Description  get single collection by id
// @Tags         collection
// @param 		 collection-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/collections/XXXX"
// @Router       /collections/{collection-id} [GET]
func GetCollectionByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  all items in collection
// @Tags         collection
// @param 		 collection-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/collections/XXXXX/items"
// @Router       /collections/{collection-id}/items [GET]
func GetAllItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  all items
// @Tags         item
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/items"
// @Router       /items [GET]
func Name11_1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  single item
// @Tags         item
// @param 		 item-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/items/yiiiiiii"
// @Router       /items/{item-id} [GET]
func GetItemById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}


