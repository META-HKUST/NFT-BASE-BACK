package v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Item struct {
	ItemName     string   `json:"item_name" example:"Pixel Bear With Hammer"`
	ItemId       string   `json:"item_id" example:"1010"`
	ItemData     string   `json:"item_data" example:"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240"`
	CreateTime   string   `json:"create_time" example:"2022-06-16 22:04:22"`
	Description  string   `json:"description" example:"A very cute pixel bear with hammer"`
	CollectionId string   `json:"collection_id" example:"Pixel Bear"`
	Category     string   `json:"category" example:"image"`
	Label        []string `json:"label" example:"Music ,Comics"`
	CreaterId    string   `json:"creater_id" example:"mingzheliu-ust-hk"`
	OwnerId      string   `json:"owner_id" example:"mingzheliu-ust-hk"`
	FavoriteNum  int64    `json:"favorite_num" example:"100"`
	Favorite     bool     `json:"favorite" example:"false"`
}

type ItemResponse struct {
	Code   int		`json:"code" example:"0"`
	Msg    string	`json:"msg" example:"Operation succeed"`
	Data   Item		`json:"data"`
}

type CreateParams struct {
	ItemData     string   `json:"item_data" example:"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240"`
	ItemName     string   `json:"item_name" example:"Pixel Bear With Hammer"`
	Description  string   `json:"description" example:"A very cute pixel bear with hammer"`
	CollectionId string   `json:"collection_id" example:"Pixel Bear"`
	Label        []string `json:"label" example:"Music ,Comics"`
}

type TransferParams struct {
	ItemId   string `json:"item_id" example:"1010"`
	ToUserId string `json:"to_user_id" example:"zhengwang-ust-hk"`
}

// CreateItem @Description  create single item: parse UserId from token and create NFT(Creater and Owner are defined as UserId)
// @Tags         item
// @param 		 item_data      formData  file  true    "NFT本身数据"
// @param 		 param_request  body  CreateParams   true   "info needed to upload"
// @Accept       json
// @Produce      json
// @Success 200 {object} ItemResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /item/create [POST]
// @Security ApiKeyAuth
func CreateItem(ctx *gin.Context) {
	resp := ItemResponse{
			0,
			"Operation succeed",
			Item{
				"Pixel Bear With Hammer",
				"1010",
				"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240",
				"2022-06-16 22:04:22",
				"A very cute pixel bear with hammer",
				"Pixel Bear",
				"image",
				[]string{"Music", "Comics"},
				"mingzheliu-ust-hk",
				"mingzheliu-ust-hk",
				100,
				false,
			},
	}
	ctx.JSON(http.StatusOK, resp)
}

type EditParams struct {
	ItemId       string   `json:"item_data" example:"1001"`
	ItemName     string   `json:"item_name" example:"Pixel Bear With Hammer"`
	Description  string   `json:"description" example:"A very cute pixel bear with hammer"`
	CollectionId string   `json:"collection_id" example:"Pixel Bear"`
	Label        []string `json:"label" example:"Music ,Comics"`
}

// EditItem @Description  edit single item
// @Tags         item
// @param 		 item_data      formData  file  false    "NFT本身数据"
// @param 		 param_request  body  EditParams  true   "info needed to upload"
// @Accept       json
// @Produce      json
// @Success 200 {object} ItemResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /item/edit [POST]
// @Security ApiKeyAuth
func EditItem(ctx *gin.Context) {
	resp := ItemResponse{
		0,
		"Operation succeed",
		Item{
			"Pixel Bear With Hammer",
			"1010",
			"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240",
			"2022-06-16 22:04:22",
			"A very cute pixel bear with hammer",
			"Pixel Bear",
			"image",
			[]string{"Music", "Comics"},
			"mingzheliu-ust-hk",
			"mingzheliu-ust-hk",
			100,
			false,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}

// TransferItem @Description  Transfer  item
// @Tags         item
// @param 		 param_request  body  TransferParams  true   "item needed to transfer"
// @Accept       json
// @Produce      json
// @Success 200 {object} ItemResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /item/transfer [POST]
// @Security ApiKeyAuth
func TransferItem(ctx *gin.Context) {
	resp := ItemResponse{
		0,
		"Operation succeed",
		Item{
			"Pixel Bear With Hammer",
			"1010",
			"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240",
			"2022-06-16 22:04:22",
			"A very cute pixel bear with hammer",
			"Pixel Bear",
			"image",
			[]string{"Music", "Comics"},
			"mingzheliu-ust-hk",
			"zhengwang-ust-hk",
			100,
			false,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}

// LikeItem @Description  edit single item
// @Tags         item
// @param 		 item_id  body string  true   "item id"
// @Accept       json
// @Produce      json
// @Success 200 {object} ItemResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /item/like [POST]
// @Security ApiKeyAuth
func LikeItem(ctx *gin.Context) {
	resp := ItemResponse{
		0,
		"Operation succeed",
		Item{
			"Pixel Bear With Hammer",
			"1010",
			"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240",
			"2022-06-16 22:04:22",
			"A very cute pixel bear with hammer",
			"Pixel Bear",
			"image",
			[]string{"Music", "Comics"},
			"mingzheliu-ust-hk",
			"zhengwang-ust-hk",
			100,
			false,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}
