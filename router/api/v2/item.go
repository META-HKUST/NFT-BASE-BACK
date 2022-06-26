package v2

import (
	"NFT-BASE-BACK/entity"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/sdk"
	"NFT-BASE-BACK/sdk/pb"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
	Code int    `json:"code" example:"0"`
	Msg  string `json:"msg" example:"Operation succeed"`
	Data Item   `json:"data"`
}

type CreateParams struct {
	ItemName string `json:"item_name" example:"Pixel Bear With Hammer"`

	ItemDataCos       string `json:"item_data_cos" example:"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240"`
	ItemDataIpfs      string `json:"item_data_ipfs" example:"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240"`
	ItemDataSignature string `json:"item_data_signature" example:"sfahelkgbekjbfbsauiehnv"`

	Description  string   `json:"description" example:"A very cute pixel bear with hammer"`
	CollectionId int      `json:"collection_id" example:"1"`
	Label        []string `json:"label" example:"Music,Comics"`

	Category string `json:"category" example:"image"`
}

type LikeRequest struct {
	ItemId string `json:"item_id" example:"1"`
}

type CreateParamsResponse struct {
	Code int         `json:"code" example:"0"`
	Msg  string      `json:"msg" example:"Operation succeed"`
	Data entity.Item `json:"data"`
}

// CreateItem @Description  create single item: parse UserId from token and create NFT(Creater and Owner are defined as UserId)
// @Tags         item
// @param 		 param_request  body  CreateParams   true   "info needed to upload"
// @Accept       json
// @Produce      json
// @Success 200 {object} CreateParamsResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /item/create [POST]
// @Security ApiKeyAuth
func CreateItem(ctx *gin.Context) {
	var req CreateParams
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// 通过jwt  email
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, "auth email error")
		return
	}
	username := strings.Replace(email.(string), "@", "-", -1)
	username = strings.Replace(username, ".", "-", -1)

	// 调用sdk
	response, err := sdk.Client.PublicMint(
		context.Background(),
		&pb.PublicMintRequest{Username: username},
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	item := model.Item{
		ItemID:       response.TokenId,
		ItemName:     req.ItemName,
		CollectionID: req.CollectionId,
		ItemData:     req.ItemDataIpfs,
		Description:  req.Description,
		OwnerID:      username,
		CreaterID:    username,
		Category:     req.Category,
	}

	// 写item数据库
	ret, err := model.CreateItem(item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "database error")
		return
	}

	// 写item label
	for _, v := range req.Label {
		itemLabel := model.ItemLable{
			ItemID:    response.TokenId,
			ItemLabel: v,
		}
		_, err = model.CreateItemLabel(itemLabel)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "database error")
			return
		}
	}
	// 在URI里面写数据
	// ？？

	// 组装response
	resp := CreateParamsResponse{
		Code: 0,
		Msg:  "Operation succeed",
		Data: entity.Item{
			ItemName:    ret.ItemName,
			ItemID:      ret.ItemID,
			ItemData:    ret.ItemData,
			CreateTime:  ret.CreatedAt,
			Description: ret.Description,

			CollectionId: ret.CollectionID,
			// CollectionName: ,
			Category: ret.Category,
			Label:    req.Label,

			CreaterId: ret.CreaterID,
			OwnerId:   ret.OwnerID,

			FavoriteNum: 0,
			Favorite:    false,
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

type TransferParams struct {
	ItemId   string `json:"item_id" example:"1010"`
	ToUserId string `json:"to_user_id" example:"zhengwang-ust-hk"`
}

type TransferItemResponse struct {
	Code int         `json:"code" example:"0"`
	Msg  string      `json:"msg" example:"Operation succeed"`
	Data entity.Item `json:"data"`
}

// TransferItem @Description  Transfer  item
// @Tags         item
// @param 		 param_request  body  TransferParams  true   "item needed to transfer"
// @Accept       json
// @Produce      json
// @Success 200 {object} TransferItemResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /item/transfer [POST]
// @Security ApiKeyAuth
func TransferItem(ctx *gin.Context) {
	var req TransferParams
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// 通过jwt，拿到email
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, "auth email error")
		return
	}
	username := strings.Replace(email.(string), "@", "-", -1)
	username = strings.Replace(username, ".", "-", -1)
	// 调用sdk
	_, err = sdk.Client.TransferFrom(
		context.Background(),
		&pb.TransferFromRequest{
			From:     username,
			To:       req.ToUserId,
			TokenId:  req.ItemId,
			Username: username,
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// 查item
	ret, err := model.UpdateItemOwner(req.ItemId, req.ToUserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// 查label
	ret_label, err := model.SearchLable(req.ItemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// 组装response
	resp := CreateParamsResponse{
		Code: 0,
		Msg:  "Operation succeed",
		Data: entity.Item{
			ItemName:    ret.ItemName,
			ItemID:      ret.ItemID,
			ItemData:    ret.ItemData,
			CreateTime:  ret.CreatedAt,
			Description: ret.Description,

			CollectionId: ret.CollectionID,
			// CollectionName: ,
			Category: ret.Category,
			Label:    ret_label,

			CreaterId: ret.CreaterID,
			OwnerId:   ret.OwnerID,

			FavoriteNum: 0,
			Favorite:    false,
		},
	}

	ctx.JSON(http.StatusOK, resp)
}

// LikeItem @Description  edit single item
// @Tags         item
// @param 		 item_id  body LikeRequest true  "item id"
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
