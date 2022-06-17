package v1

import (
	"NFT-BASE-BACK/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)


type NewItem struct {
	// basic info
	ItemName		string   	`josn:"item_name" example:"imageNFT"`
	ItemId			string      `json:"item_id" example:"123455"`
	ItemData		string   	`json:"item_data" exmaple:"http://www.image.com/123455"`
	CreateTime		string   	`json:"create-time"`
	Description		string   	`json:"description"`
	// characteristic
	CollectionID	string   	`json:"collection-id" example:"5"`
	Category		string   	`json:"category"`
	Label			[]string 	`json:"lable"`
	// account id info
	CreaterId		string 		`json:"creater-id" example:"mazhengwang-ust-hk"`
	OwnerId			string 		`json:"owner-id" example:"mazhengwang-ust-hk"`
	FavoriteNum		int 		`json:"favorite_num"`
	Favorite		bool		`json:"favortie"`
}

type ItemResponse struct {
	Code			int		`json:"code"`
	Msg 			string		`json:"msg"`
	Data 			NewItem
}

// CreateItem @Description  create single item: parse UserId from token and create NFT(Creater and Owner are defined as UserId)
// @Tags         item
// @param 		 item_name   query   string   true   "item_name"
// @param 		 item_data   query   string   true    "item_data"
// @param 		 description   query   string   true    "description"
// @param 		 item_collection   query   string   true    "item_collection"
// @param 		 label   query   string   true    "label"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 1000 {object} base.ErrCode "request error"
// @Failure 2000 {object} base.Response "error code and message and nil data"
// @Router       /item/create [POST]
func CreateItem(ctx *gin.Context) {

	name := ctx.Query("item_name")
	data := ctx.Query("item_data")
	description := ctx.Query("description")
	collection := ctx.Query("item_collection")
	label := ctx.Query("label")

	newitem := NewItem{
		name,
		"1001",
		data,
		utils.GetTimeNow(),
		description,
		collection,
		"image",
		[]string{label},
		"mingzheliu-ust-hk",
		"mingzheliu-ust-hk",
		100,
		false,
	}
	resp := ItemResponse{
		0,
		"Operation succeed",
		newitem,
	}
	ctx.JSON(http.StatusOK, resp)
}

// EditItem @Description  edit single item: parse UserId from token and edit NFT
// @Tags         item
// @param 		 item-id   query   string   false   "item-id"
// @param 		 item_name   query   string   false   "name"
// @param 		 description   query   string   false   "description"
// @param 		 collection_id   query   string   false   "collection"
// @param 		 label   query   string   false   "label"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 1000 {object} base.ErrCode "request error"
// @Failure 2000 {object} base.Response "error code and message and nil data"
// @Router       /item/edit [POST]
func EditItem(ctx *gin.Context) {


	name := ctx.Query("item_name")
	item_id := ctx.Query("item-id")

	description := ctx.Query("description")
	itemCollection := ctx.Query("collection_id")
	label := ctx.Query("label")
	image := "/yezzi/1.png"
	category := "image"

	item := NewItem{
		name,
		item_id,
		image,
		utils.GetTimeNow(),
		description,
		itemCollection,
		category,
		[]string{label},
		"mingzheliu-ust-hk",
		"mingzheliu-ust-hk",
		100,
		false,

	}
	resp :=ItemResponse{
		0,
		"Operation succeed",
		item,
	}
	ctx.JSON(http.StatusOK, resp)
}


// EditItem @Description  edit single item: parse UserId from token and edit NFT
// @Tags         item
// @param 		 to_user_id   query   string   false   "to user id"
// @param 		 item_id   query   string   false   "item id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed"
// @Failure 1000 {object} base.ErrCode "request error"
// @Failure 2000 {object} base.Response "error code and message and nil data"
// @Router       /item/transfer [POST]
func TransferItem(ctx *gin.Context) {
	toUserId := ctx.Query("to_user_id")
	ItemId := ctx.Query("item_id")

	item := NewItem{
		"Pixel Bear With Hammer",
		ItemId,
		"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240",
		utils.GetTimeNow(),
		"A very cute pixel bear with hammer",
		"Pixel Bear",
		"image",
		[]string{"pixel","bear"},
		"mingzheliu-ust-hk",
		toUserId,
		100,
		false,

	}
	resp :=ItemResponse{
		0,
		"Operation succeed",
		item,
	}
	ctx.JSON(http.StatusOK, resp)
}


// EditItem @Description  edit single item: parse UserId from token and edit NFT
// @Tags         item
// @param 		 item-id   query   string   false   "item-id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed"
// @Failure 1000 {object} base.ErrCode "request error"
// @Failure 2000 {object} base.Response "error code and message and nil data"
// @Router       /item/like [POST]
func LikeItem(ctx *gin.Context)  {

	ItemId := ctx.Query("item_id")

	item := NewItem{
		"Pixel Bear With Hammer",
		ItemId,
		"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240",
		utils.GetTimeNow(),
		"A very cute pixel bear with hammer",
		"Pixel Bear",
		"image",
		[]string{"pixel","bear"},
		"mingzheliu-ust-hk",
		"mingzheliu-ust-hk",
		100,
		true,

	}
	resp :=ItemResponse{
		0,
		"Operation succeed",
		item,
	}
	ctx.JSON(http.StatusOK, resp)
}


