package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserID            string `json:"user_id" example:"ingzheliu-ust-hk"`
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

	users := []UserInfo{
		{"mingzheliu-ust-hk",
			"mingzheliu@ust.hk",
			"LMZ",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
			"teacher",
			"HKUST-GZ",
			"2022-06-16 20:45:40"},
		{
			"mingzheliu-ust-hk",
			"mingzheliu@ust.hk",
			"LMZ",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
			"teacher",
			"HKUST-GZ",
			"2022-06-16 20:45:40",
		},
	}
	resp := ListResponse{
		0,
		"Operation succeed",
		UsersList{
			users,
			1,
			10,
			1,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}

// SingleColletction @Description  get all users in database
// @Tags         list
// @param 		 collection_id   query   string   true   "collection id"
// @Accept       json
// @Produce      json
// @Success 200 {object} ListResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/collection [GET]
func SingleColletction(ctx *gin.Context) {
	res := base.Response{}
	var ch int
	ctx.BindJSON(&ch)
	data, err := service.GetCollection(ch)
	if err != nil {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
	}
	res.SetData(data)

	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

// CollectionList @Description  get all users in database
// @Tags         list
// @param 		 keyword   path   string   false   "keyword"
// @param 		 rank_favorite   path   bool   false   "rank favorite"
// @param 		 rank_time   path   bool   false   "rank time"
// @param 		 user_id   path   string   false   "user id"
// @param 		 label   path   string   false   "label"
// @param 		 page_num   path   int   false   "page num"
// @param 		 page_size   path   int   false   "page size"
// @Accept       json
// @Produce      json
// @Success 200 {object} ListResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/collection-list [GET]
func CollectionList(ctx *gin.Context) {
	resp := ListResponse{
		0,
		"Operation succeed",
		CollectionsList{
			[]CollectionInfo{
				{
					"Doodles",
					"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
					"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
					"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
					"A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Commun",
					[]string{"Comics"},
					20,
					"ZZD",
					"2022-06-16 20:45:40",
				},
				{
					"MAYC",
					"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
					"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
					"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
					"The MUTANT APE YACHT CLUB is a collection of up to 20,000 Mutant Apes that can only be created by exposing an existing Bored Ape to a vial of MUTANT SERUM or by minting a Mutant Ape in the public sale.",
					[]string{"Comics"},
					20,
					"ZW",
					"2022-06-16 20:45:40",
				},
			},
			1,
			10,
			1,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}

// SingleItem @Description  get all users in database
// @Tags         list
// @param 		 item_id   query   string   true   "item id"
// @Accept       json
// @Produce      json
// @Success 200 {object} ListResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/item [GET]
func SingleItem(ctx *gin.Context) {
	var resp base.Response
	itemId, ok := ctx.Get("item_id")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, "Input parameter error")
		return
	}

	itemInfo, code := service.GetItem(itemId.(string))
	if code != base.Success {
		ctx.JSON(http.StatusInternalServerError, "Failed to get items information")
		return
	}

	resp.SetData(itemInfo)
	resp.SetCode(code)
	ctx.JSON(http.StatusOK, resp)
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
	_, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, "auth email error")
		return
	}

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

	items, code := service.GetItemList(pageNumInt, pageSizeInt, userId, userLikeBool, userCollectBool, userCreateBool, category, keyword, rankFavoriteBool, rankTimeBool, collectionIdInt)
	if code != nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to get items information")
		return
	}
	resp.SetData(items)
	ctx.JSON(http.StatusOK, resp.SetCode(0))
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
	var req ItemHistoryRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	hs, err := model.GetItemHistory(req.ItemID)
	if err != nil {
		ctx.JSON(http.StatusOK, base.ServerError)
	}
	resp.SetCode(base.Success)
	resp.SetData(hs)
	ctx.JSON(http.StatusOK, resp)

}
