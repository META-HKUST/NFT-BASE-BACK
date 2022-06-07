package v1

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllCollections @Description  get all collections in database using some methods
// @Tags         collection
// @param 		 num   query   string   true   "num"
// @param 		 pagesize   query   string   true   "pagesize"
// @param 		 method   query   string   true   "method on how to sort these collections"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /collections [GET]
func GetAllCollections(ctx *gin.Context) {
	res := base.PageResponse{}
	pgnumber := ctx.Query("num")
	pgsize := ctx.Query("pagesize")
	methond := ctx.Query("method")
	fmt.Println(pgnumber, pgsize)
	code, data, count := service.GetCollections("", 1, 2, methond)
	res.SetCode(code)
	res.SetData(data)
	res.SetCount(count)
	ctx.JSON(http.StatusOK, res)
}

// GetCollectionByID @Description  get single collection by id
// @Tags         collection
// @param 		 collection-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /collections/{collection-id} [GET]
func GetCollectionByID(ctx *gin.Context) {
	res := base.Response{}
	CollectionId := ctx.Query("collection-id")
	code, data := service.GetCollection(CollectionId)
	res.SetData(data)
	res.SetCode(code)
	ctx.JSON(http.StatusOK, res)
}
