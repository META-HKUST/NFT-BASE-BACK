package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/entity"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/sdk"
	"NFT-BASE-BACK/sdk/pb"
	"NFT-BASE-BACK/service"
	"NFT-BASE-BACK/utils"
	"context"
	"fmt"
	"log"
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
	FavoriteNum  int64    `json:"like_count" example:"100"`
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
	CollectionId string   `json:"collection_id" example:"1"`
	Label        []string `json:"label" example:"Music,Comics"`

	Category string `json:"category" example:"image"`
}

type CreateParamsResponse struct {
	Code int         `json:"code" example:"0"`
	Msg  string      `json:"msg" example:"Operation succeed"`
	Data entity.Item `json:"data"`
}

type ItemRes struct {
	Item            interface{} `json:"item" example:"0"`
	Collection_name string      `json:"collection_name" example:"0"`
	Owner_name      string      `json:"owner_name" example:"0"`
	Creater_name    string      `json:"creater_name" example:"0"`
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

	// check empty
	ss := append([]string{}, req.ItemName, req.ItemDataCos, req.ItemDataIpfs, req.ItemDataSignature, req.Description, req.Category)
	if utils.CheckAnyEmpty(ss) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}
	// check empty
	if utils.CheckEmpty(req.CollectionId) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}

	// 通过jwt  email
	email, ok := ctx.Get("email")
	if !ok {
		res := base.Response{}
		ctx.JSON(http.StatusOK, res.SetCode(base.InputError))
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
		log.Println(err)
		res := base.Response{}
		ctx.JSON(http.StatusOK, res.SetCode(base.FabricInvokeError))
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
	//item := model.Item{
	//	ItemID:       utils.GenEmailToken(),
	//	ItemName:     "aaa",
	//	CollectionID: 1,
	//	ItemData:     "aaa",
	//	Description:  "aaa",
	//	OwnerID:      "aaa",
	//	CreaterID:    "aaa",
	//	Category:     "aaa",
	//}

	// 写item数据库
	ret, err := model.CreateItem(item)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, "database error")
		return
	}

	// TODO：根据tokenURI 去构造IPFS相关的json，存数据库里，之后暴露这个接口能让人查到

	// 写item label
	for _, v := range req.Label {
		itemLabel := model.ItemLable{
			// change this to test
			ItemID:    item.ItemID,
			ItemLabel: v,
		}
		_, err = model.CreateItemLabel(itemLabel)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, "database error")
			return
		}
	}
	// 在URI里面写数据
	// ？？

	baseItem := entity.Item{
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
	}
	collectionName, _ := model.GetCollectionName(baseItem.CollectionId)
	ownerName, _ := model.GetUserName(baseItem.OwnerId)
	createrName, _ := model.GetUserName(baseItem.CreaterId)
	resData := ItemRes{
		baseItem,
		collectionName,
		ownerName,
		createrName,
	}

	err = model.AddTransferHistory(ret.ItemID, "Create", username)
	if err != nil {
		log.Println(err)
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.ServerError))
	}

	// 组装response
	resp := base.Response{}
	resp.SetCode(base.Success)
	resp.SetData(resData)

	ctx.JSON(http.StatusOK, resp)
}

type EditParams struct {
	ItemId       string   `json:"item_id" example:"1001"`
	ItemName     string   `json:"item_name" example:"Pixel Bear With Hammer"`
	Description  string   `json:"description" example:"A very cute pixel bear with hammer"`
	CollectionId string   `json:"collection_id" example:"Pixel Bear"`
	Label        []string `json:"label" example:"Music ,Comics"`
}

type UpdateParams struct {
	TokenId string `json:"token_id" example:"1001"`
	IpfsUrl string `json:"ipfs_url" example:"http://ipfs/xmaedhkdhfrfndj"`
}

// EditItem @Description  edit single item
// @Tags         item
// @param 		 param_request  body  EditParams  true   "info needed to upload"
// @Accept       json
// @Produce      json
// @Success 200 {object} ItemResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /item/edit [POST]
// @Security ApiKeyAuth
func EditItem(ctx *gin.Context) {
	var resp base.Response
	var req EditParams

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	// check empty
	if utils.CheckEmpty(req.ItemId) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}

	itemInfo, code := service.EditItem(req.ItemId, req.ItemName, req.Description, req.CollectionId, req.Label)
	if code != base.Success {
		ctx.JSON(http.StatusBadRequest, resp.SetCode(code))
		return
	}
	collectionName, _ := model.GetCollectionName(itemInfo.CollectionID)
	ownerName, _ := model.GetUserName(itemInfo.OwnerID)
	createrName, _ := model.GetUserName(itemInfo.CreaterID)
	resData := ItemRes{
		itemInfo,
		collectionName,
		ownerName,
		createrName,
	}
	resp.SetCode(code)
	resp.SetData(resData)
	ctx.JSON(http.StatusOK, resp)
}

// UpdateItem @Description  edit single item
// @Tags         item
// @param 		 param_request  body  UpdateParams  true   "info needed to update"
// @Accept       json
// @Produce      json
// @Success 200 {object} ItemResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /item/update [POST]
// @Security ApiKeyAuth
func UpdateItem(ctx *gin.Context) {
	var resp base.Response
	var req UpdateParams

	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusOK, new(base.Response).SetCode(base.AuthFailed))
		return
	}
	if email != "admin@unifit.art" {

		ctx.JSON(http.StatusOK, new(base.Response).SetCode(base.AuthFailed))
		return
	}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	itemInfo, code := service.UpdateItem(req.TokenId, req.IpfsUrl)
	if code != base.Success {
		ctx.JSON(http.StatusBadRequest, resp.SetCode(code))
		return
	}

	collectionName, _ := model.GetCollectionName(itemInfo.CollectionID)
	ownerName, _ := model.GetUserName(itemInfo.OwnerID)
	createrName, _ := model.GetUserName(itemInfo.CreaterID)
	resData := ItemRes{
		itemInfo,
		collectionName,
		ownerName,
		createrName,
	}
	resp.SetCode(code)
	resp.SetData(resData)
	ctx.JSON(http.StatusOK, resp)
}

type TransferParams struct {
	ItemId  string `json:"item_id" example:"1010"`
	ToEmail string `json:"email" example:"zhengwang@ust.hk"`
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
		ctx.JSON(http.StatusBadRequest, new(base.Response).SetCode(base.InputError))
		return
	}

	email2 := req.ToEmail
	username2 := strings.Replace(email2, "@", "-", -1)
	username2 = strings.Replace(username2, ".", "-", -1)

	// check empty
	ss := append([]string{}, req.ItemId, username2)
	if utils.CheckAnyEmpty(ss) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}

	// 通过jwt，拿到email
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusOK, new(base.Response).SetCode(base.AuthFailed))
		return
	}
	username := strings.Replace(email.(string), "@", "-", -1)
	username = strings.Replace(username, ".", "-", -1)

	// 调用sdk
	_, err = sdk.Client.TransferFrom(
		context.Background(),
		&pb.TransferFromRequest{
			From:     username,
			To:       username2,
			TokenId:  req.ItemId,
			Username: username,
		},
	)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, new(base.Response).SetCode(base.FabricInvokeError))
		return
	}

	// 查item
	ret, err := model.UpdateItemOwner(req.ItemId, username2)
	if err != nil {
		ctx.JSON(http.StatusOK, new(base.Response).SetCode(base.ServerError))
		return
	}

	// 查label
	ret_label, err := model.SearchLable(req.ItemId)
	if err != nil {
		ctx.JSON(http.StatusOK, new(base.Response).SetCode(base.ServerError))
		return
	}

	likeCount, _ := model.GetLikeCount(ret.ItemID)
	baseItem := entity.Item{
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

		FavoriteNum: likeCount,
		Favorite:    false,
	}

	collectionName, _ := model.GetCollectionName(baseItem.CollectionId)
	ownerName, _ := model.GetUserName(baseItem.OwnerId)
	createrName, _ := model.GetUserName(baseItem.CreaterId)
	resData := ItemRes{
		baseItem,
		collectionName,
		ownerName,
		createrName,
	}
	// 组装response
	resp := base.Response{}
	resp.SetCode(base.Success)
	resp.SetData(resData)

	// add item history
	err = model.AddTransferHistory(ret.ItemID, username, username2)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, resp.SetCode(base.ServerError))
	}

	ctx.JSON(http.StatusOK, resp)
}

type LikeRequest struct {
	ItemId string `json:"item_id" example:"1"`
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
	ch := PostActUploadItemRequest{}
	ctx.BindJSON(&ch)
	// check empty
	if utils.CheckEmpty(ch.ItemID) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}
	res := base.Response{}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)
	fmt.Println(UserId)

	bo, _ := model.DoesLike(ch.ItemID, UserId)
	log.Println("Invoke like item, status: ", UserId, " ", ch.ItemID, " ", bo)
	if bo == false {
		log.Println("User ", UserId, " will like this item:", ch.ItemID)
		err := model.Like(ch.ItemID, UserId)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
			return
		}
		ctx.JSON(http.StatusOK, res.SetCode(base.Success))
	} else {
		log.Println("User ", UserId, " will unlike this item:", ch.ItemID)
		err := model.UnLike(ch.ItemID, UserId)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
			return
		}
		ctx.JSON(http.StatusOK, res.SetCode(base.Success))
	}

}
