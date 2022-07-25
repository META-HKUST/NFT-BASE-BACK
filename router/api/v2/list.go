package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/service"
	"NFT-BASE-BACK/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserID            string `json:"user_id" example:"mingzheliu-ust-hk"`
	UserEmail         string `json:"user_email" example:"mingzheliu@ust.hk"`
	UserName          string `json:"user_name" example:"LMZ"`
	BannerImage       string `json:"banner_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	LogoImage         string `json:"logo_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e"`
	Poison            string `json:"poison" example:"teacher"`
	Organization      string `json:"organization" example:"HKUST-GZ"`
	RegisterationTime string `json:"registeration_time" example:"2022-06-16 20:45:40"`
}

type ListResponse struct {
	Code int         `json:"code" example:"0"`
	Msg  string      `json:"msg" example:"Operation succeed"`
	Data interface{} `json:"data"`
}

type UsersList struct {
	UserList interface{} `json:"user_list"`
	Page     int         `json:"page"`
	Size     int         `json:"size"`
	Total    int         `json:"total"`
}

type CollectionsList struct {
	CollectionList interface{} `json:"collection_list"`
	Page           int         `json:"page"`
	Size           int         `json:"size"`
	Total          int         `json:"total"`
}

type ItemsList struct {
	Data  interface{} `json:"data"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
}

type Collection struct {
	CollectionId   string   `json:"collection_id" example:"111111"`
	CollectionName string   `json:"collection_name" example:"Doodles"`
	BannerImage    string   `json:"banner_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	LogoImage      string   `json:"logo_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e"`
	FeatureImage   string   `json:"feature_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	Description    string   `json:"description" example:"A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Commun"`
	Label          []string `json:"label" example:"Comics"`
	ItemNum        int      `json:"item_num" example:"20"`
	OwnerId        string   `json:"owner_id" example:"zezhending-ust-hk"`
	OwnerName      string   `json:"owner_name" example:"ZZD"`
	CreationTime   string   `json:"creation_time" example:"2022-06-16 20:45:40"`
}

type CollectionInfo struct {
	CollectionName string   `json:"collection_name" example:"Doodles"`
	BannerImage    string   `json:"banner_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	LogoImage      string   `json:"logo_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e"`
	FeatureImage   string   `json:"feature_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	Description    string   `json:"description" example:"A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Commun"`
	Label          []string `json:"label" example:"Comics"`
	ItemNum        int      `json:"item_num" example:"20"`
	OwnerName      string   `json:"owner_name" example:"ZZD"`
	CreationTime   string   `json:"creation_time" example:"2022-06-16 20:45:40"`
}

type History struct {
	FromUserId   string `json:"from_user_id"`
	ToUserId     string `json:"to_user_id"`
	FromUserName string `json:"from_user_name"`
	ToUserName   string `json:"to_user_name"`
	Time         string `json:"time"`
}

// UserList @Description  get all users in database
// @Tags         list
// @param 		 keyword   query   string   false   "keyword"
// @param 		 page_size  query  int   true   "pagesize"
// @param 		 page_num   query  int   true   "page num"
// @Accept       json
// @Produce      json
// @Success 200 {object} ListResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/user-list [GET]
func UserList(ctx *gin.Context) {
	var resp base.Response
	pageNum := ctx.Query("page_num")
	pageNumInt, _ := strconv.ParseInt(pageNum, 10, 64)
	pageSize := ctx.Query("page_size")
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)
	keyword := ctx.Query("keyword")

	userList, err := service.GetUserListByKeyWord(keyword, pageNumInt, pageSizeInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get user informations")
		return
	}
	resp.SetData(userList)
	ctx.JSON(http.StatusOK, resp.SetCode(0))
}

// SingleColletction @Description  get all users in database
// @Tags         list
// @param 		 collection_id   query   string   true   "collection id"
// @Accept       json
// @Produce      json
// @Success 200 {object} Collection "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/collection [GET]
func SingleColletction(ctx *gin.Context) {
	res := base.Response{}
	ch := ctx.Query("collection_id")

	// check empty
	if utils.CheckEmpty(ch) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}

	data, err := model.GetCoAndLabel(ch)

	if err != nil {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(data)

	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type ColistRes struct {
	Colist []model.Collection
	Count  int `json:"count"`
}

// CollectionList @Description  get all users in database
// @Tags         list
// @param 		 page_num   query   int   true   "page num"
// @param 		 page_size   query   int   true   "page size"
// @param 		 keyword   query   string   false   "keyword"
// @param 		 rank_favorite   query   bool   false   "rank favorite"
// @param 		 rank_time   query   bool   false   "rank time"
// @param 		 user_id   query   string   false   "user id"
// @param 		 label   query   string   false   "label"
// @Accept       json
// @Produce      json
// @Success 200 {object} ListResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/collection-list [GET]
func CollectionList(ctx *gin.Context) {

	var resp base.Response

	pageNum := ctx.Query("page_num")
	pageNumInt, _ := strconv.ParseInt(pageNum, 10, 64)
	pageSize := ctx.Query("page_size")
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)
	label := ctx.Query("label")
	keyword := ctx.Query("keyword")
	rankFavorite := ctx.Query("rank_favorite")
	rankFavoriteBool, _ := strconv.ParseBool(rankFavorite)
	rankTime := ctx.Query("rank_time")
	rankTimeBool, _ := strconv.ParseBool(rankTime)
	userId := ctx.Query("user_id")

	var coRes ColistRes
	var err error
	//collections, code := service.GetCollectionList(pageNumInt, pageSizeInt, userId, keyword, rankFavoriteBool, rankTimeBool, label)
	coRes.Colist, coRes.Count, err = model.GetCollectionList(pageNumInt, pageSizeInt, userId, keyword, rankFavoriteBool, rankTimeBool, label)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, resp.SetCode(base.ServerError))
	}

	resp.SetData(coRes)
	ctx.JSON(http.StatusOK, resp.SetCode(base.Success))
}

// SingleItem @Description  get all users in database
// @Tags         list
// @param 		 item_id   query   string   true   "item id"
// @Accept       json
// @Produce      json
// @Success 200 {object} ItemResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/item [GET]
func SingleItem(ctx *gin.Context) {
	var resp base.Response
	itemId := ctx.Query("item_id")

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	//// check if admin account
	//if email != "1721062927@qq.com" {
	//	ctx.JSON(http.StatusOK, base.PermissionDenied)
	//}

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)
	log.Println("user_id: ", UserId)

	baseItem, code := service.GetItem(itemId, UserId)

	if code != base.Success {
		ctx.JSON(http.StatusInternalServerError, "Failed to get items information")
		return
	}

	collectionName, _ := model.GetCollectionName(baseItem.CollectionID)
	ownerName, _ := model.GetUserName(baseItem.OwnerID)
	createrName, _ := model.GetUserName(baseItem.CreaterID)
	resData := ItemRes{
		baseItem,
		collectionName,
		ownerName,
		createrName,
	}
	// 组装response
	resp.SetCode(base.Success)
	resp.SetData(resData)

	ctx.JSON(http.StatusOK, resp)
}

type ItemlistRes struct {
	ItemALs []model.ItemAndLogo
	Count   int `json:"count"`
}

// ItemList @Description  get all users in database
// @Tags         list
// @param 		 page_num   query   int   true   "page num"
// @param 		 page_size  query   int   true   "page size"
// @param 		 user_id   query   string   false   "user id"
// @param 		 user_like  query   bool   false   "user like"
// @param 		 user_collect   query   bool   false   "user collect"
// @param 		 user_create   query   bool   false   "user create"
// @param 		 category   query   string   false   "category"
// @param 		 keyword   query   string   false   "keyword"
// @param 		 rank_favorite   query   bool   false   "rank favorite"
// @param 		 rank_time   query   bool   false   "rank time"
// @param 		 collection_id   query   string   false   "collection id"
// @param 		 label   query   string   false   "label"
// @Accept       json
// @Produce      json
// @Success 200 {object} ListResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/item-list [GET]
func ItemList(ctx *gin.Context) {
	var resp base.Response

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	//// check if admin account
	//if email != "1721062927@qq.com" {
	//	ctx.JSON(http.StatusOK, base.PermissionDenied)
	//}

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)

	pageNum := ctx.Query("page_num")
	pageNumInt, _ := strconv.ParseInt(pageNum, 10, 64)
	pageSize := ctx.Query("page_size")
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)
	userId := ctx.Query("user_id")
	userLike := ctx.Query("user_like")
	userLikeBool, _ := strconv.ParseBool(userLike)
	userCollect := ctx.Query("user_collect")
	userCollectBool, _ := strconv.ParseBool(userCollect)
	userCreate := ctx.Query("user_create")
	userCreateBool, _ := strconv.ParseBool(userCreate)
	category := ctx.Query("category")
	keyword := ctx.Query("keyword")
	rankFavorite := ctx.Query("rank_favorite")
	rankFavoriteBool, _ := strconv.ParseBool(rankFavorite)
	rankTime := ctx.Query("rank_time")
	rankTimeBool, _ := strconv.ParseBool(rankTime)
	collectionId := ctx.Query("collection_id")
	collectionIdInt, _ := strconv.Atoi(collectionId)

	if userId != "" {
		log.Println("item-list using user_id: ", userId)
	} else if UserId != "" {
		userId = UserId
		log.Println("item-list using user_id: ", userId)
	}

	var ItemsRes ItemlistRes
	var err error
	ItemsRes.ItemALs, ItemsRes.Count, err = model.GetItemList(pageNumInt, pageSizeInt, userId, userLikeBool, userCollectBool, userCreateBool, category, keyword, rankFavoriteBool, rankTimeBool, collectionIdInt)

	//items, code := service.GetItemList(pageNumInt, pageSizeInt, userId, userLikeBool, userCollectBool, userCreateBool, category, keyword, rankFavoriteBool, rankTimeBool, collectionIdInt)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.SetCode(base.ServerError))
		return
	}
	resp.SetData(ItemsRes)
	ctx.JSON(http.StatusOK, resp.SetCode(base.Success))
}

type ItemHistoryRequest struct {
	PageNum  int64  `json:"page_num" example:"1"`
	PageSize int64  `json:"page_size" example:"1"`
	ItemID   string `json:"item_id" example:"2"`
}

// ItemHistory @Description  get all users in database
// @Tags         list
// @param 		 item_id   query   string   true   "item id"
// @Accept       json
// @Produce      json
// @Success 200 {object} ListResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/item-history [GET]
func ItemHistory(ctx *gin.Context) {
	var resp base.Response

	item_id := ctx.Query("item_id")
	// check empty
	if utils.CheckEmpty(item_id) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}
	data, err := model.GetItemHistory(item_id)
	if err != nil {
		ctx.JSON(http.StatusOK, base.ServerError)
		return
	}
	resp.SetCode(base.Success)
	resp.SetData(data)
	ctx.JSON(http.StatusOK, resp)
}
