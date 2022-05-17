package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AllItemsReq struct {
	SortBy string `json:"sort-by" example:"newest"`
	Filter string `json:"filter" example:"image"`
}

type Item struct {
	ItemID         int    `json:"item_id" example:"123455"`
	ItemCollection int    `json:"collection_id" example:"5"`
	OwnerId        string `josn:"owner_id" example:"mazhengwang-ust-hk"`
	CreaterId      string `josn:"creater_id" example:"mazhengwang-ust-hk"`
	Image          string `json:"image" exmaple:"http://www.iamge.com/123455"`
	Favorites      int    `josn:"favorites" example:"1"`
}

type AllItemsResponse struct {
	Code  string `json:"code" example:"success"`
	Total int    `json:"total" example:"10"`
	Items []Item
}

// @Description  all items
// @Tags         item
// @param        sort-by  query string  false  "favorites,popular,newest"
// @param        filter   query string  false  "image,video,audio"
// @Accept       json
// @Produce      json
// @Success      200  {object}  AllItemsResponse "GET/api/v1/items"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /items [GET]
func AllItems(ctx *gin.Context) {
	fmt.Println("进入")
	var resp AllItemsResponse
	resp.Code = "SUCCESS"
	for i := 0; i < 10; i++ {
		resp.Total += 1
		ItemTmp := Item{
			ItemID:         i,
			ItemCollection: 1,
			OwnerId:        "mazhengwang-ust-hk",
			Image:          "www.image.com/" + strconv.Itoa(i),
			Favorites:      i,
		}
		resp.Items = append(resp.Items, ItemTmp)
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Description  single item
// @Tags         item
// @param 		 item-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  Item       "GET/api/v1/items/yiiiiiii"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /items/{item-id} [GET]
func SingleItem(ctx *gin.Context) {
	resp := Item{
		ItemID:         1,
		ItemCollection: 1,
		OwnerId:        "mazhengwang-ust-hk",
		Image:          "www.image.com/1",
		Favorites:      1,
	}
	ctx.JSON(http.StatusOK, resp)
}

