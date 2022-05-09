package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//// 输出数据结构体
//type P struct {
//	// HTTP 状态码 & 自定义状态码
//	Code int
//	// 输出消息
//	Message string
//	// 输出自定义数据
//	Params gin.H
//}

// @Description  get user:if username == id, return all information includes password
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna"
// @Router       /{id} [GET]
func GetUser(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  collected
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna/collected"
// @Router       /{id}/collected [GET]
func Collected(ctx *gin.Context) {
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
func Favorites(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  creation
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "GET/api/v1/Anna/creation"
// @Router       /{id}/creation [GET]
func Creation(ctx *gin.Context) {
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
func CreateItem(ctx *gin.Context) {
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
func DeleteItem(ctx *gin.Context) {
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
func CreateCollectionByAccount(ctx *gin.Context) {
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
func EditCollection(ctx *gin.Context) {
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

func EditItem(ctx *gin.Context) {
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
func ChangeProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}
