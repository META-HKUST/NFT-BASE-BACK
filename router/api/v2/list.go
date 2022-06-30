package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/service"
	"net/http"

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
	ItemList interface{} `json:"item_list"`
	Page     int         `json:"page"`
	Size     int         `json:"size"`
	Total    int         `json:"total"`
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
// @Security ApiKeyAuth
func SingleItem(ctx *gin.Context) {
	resp := ListResponse{
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

// ItemList @Description  get all users in database
// @Tags         list
// @param 		 page_num   path   int   true   "page num"
// @param 		 page_size   path   int   true   "page size"
// @param 		 user_id   path   string   false   "user id"
// @param 		 user_like   path   bool   false   "user like"
// @param 		 user_collect   path   bool   false   "user collect"
// @param 		 user_create   path   bool   false   "user create"
// @param 		 category   path   string   false   "category"
// @param 		 keyword   path   string   false   "keyword"
// @param 		 rank_favorite   path   bool   false   "rank favorite"
// @param 		 rank_time   path   bool   false   "rank time"
// @param 		 collection_id   path   string   false   "collection id"
// @param 		 label   path   string   false   "label"
// @Accept       json
// @Produce      json
// @Success 200 {object} ListResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /list/item-list [GET]
// @Security ApiKeyAuth
func ItemList(ctx *gin.Context) {
	resp := ListResponse{
		0,
		"Operation succeed",
		ItemsList{
			ItemList: []Item{
				{
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
				{
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
			},
			Page:  1,
			Size:  10,
			Total: 1,
		},
	}
	ctx.JSON(http.StatusOK, resp)
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
	resp := ListResponse{
		0,
		"Operation succeed",
		[]History{
			{
				"baofuhan-ust-hk",
				"zwang-ust-hk",
				"BFH",
				"ZW",
				"2022-06-16 20:45:40",
			},
			{
				"baofuhan-ust-hk",
				"zwang-ust-hk",
				"BFH",
				"ZW",
				"2022-06-16 20:45:40",
			},
		},
	}
	ctx.JSON(http.StatusOK, resp)
}
