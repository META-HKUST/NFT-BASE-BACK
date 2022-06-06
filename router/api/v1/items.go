package v1

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description  get all items in database and get them sorted according to "method", default time
// @Tags         item
// @param 		 pagenumber   query   string   true   "pagenumber"
// @param 		 pagesize   query   string   true   "pagesize"
// @param 		 method   query   string   true   "method on how to sort these items"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /items [GET]
func SortedItems(ctx *gin.Context) {
	res := base.PageResponse{}
	pgnumber := ctx.Query("pagenumber")
	pgsize := ctx.Query("pagesize")
	method := ctx.Query("method")
	fmt.Println(pgnumber, pgsize)
	// 如果id为空，则查询为所有items，如果id为UserId则查询一个人创建的所有items，如果id为CollectionId，则返回一个collection的所有items
	// method为排序方式
	code, data, count := service.GetItems("", 1, 2, method)
	res.SetCode(code)
	res.SetData(data)
	res.SetCount(count)
	ctx.JSON(http.StatusOK, res)
}

// @Description  single item according to item-id
// @Tags         item
// @param 		 item-id   path   string    true    "item-id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /items/{item-id} [GET]
func SingleItem(ctx *gin.Context) {
	res := base.Response{}
	ItemId := ctx.Query("item-id")
	code, data := service.GetItem(ItemId)
	res.SetData(data)
	res.SetCode(code)
	ctx.JSON(http.StatusOK, res)
}
