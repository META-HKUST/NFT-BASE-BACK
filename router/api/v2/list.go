package v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserID				string		`json:"user_id" example:"ingzheliu-ust-hk"`
	UserEmail 			string		`json:"user_email" example:"mingzheliu@ust.hk"`
	UserName			string		`json:"user_name" example:"LMZ"`
	BannerImage			string		`json:"banner_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	LogoImage			string		`json:"logo_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e"`
	Poison				string		`json:"poison" example:"teacher"`
	Organization		string		`json:"organization" example:"HKUST-GZ"`
	RegisterationTime	string		`json:"registeration_time" example:"2022-06-16 20:45:40"`
}

type ListResponse struct {
	Status 	StatusInfo
	Data 	interface{}
}

type InfoList struct {
	Info 	interface{}
	page 	int
	size 	int
	total 	int
}

type Collection struct {
	CollectionId	string		`json:"collection_id" example:"111111"`
	CollectionName	string		`json:"collection_name" example:"Doodles"`
	BannerImage		string		`json:"banner_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	LogoImage		string		`json:"logo_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e"`
	FeatureImage	string		`json:"feature_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	Description		string		`json:"description" example:"A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Commun"`
	Label			[]string	`json:"label" example:"Comics"`
	ItemNum			int			`json:"item_num" example:"20"`
	OwnerId			string		`json:"owner_id" example:"zezhending-ust-hk"`
	OwnerName		string		`json:"owner_name" example:"ZZD"`
	CreationTime	string		`json:"creation_time" example:"2022-06-16 20:45:40"`
}

type CollectionInfo struct {
	CollectionName	string		`json:"collection_name" example:"Doodles"`
	BannerImage		string		`json:"banner_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	LogoImage		string		`json:"logo_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e"`
	FeatureImage	string		`json:"feature_image" example:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	Description		string		`json:"description" example:"A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Commun"`
	Label			[]string	`json:"label" example:"Comics"`
	ItemNum			int			`json:"item_num" example:"20"`
	OwnerName		string		`json:"owner_name" example:"ZZD"`
	CreationTime	string		`json:"creation_time" example:"2022-06-16 20:45:40"`
}


type History struct {
	FromUserId 		string
	ToUserId		string
	FromUserName	string
	ToUserName		string
	Time			string
}

// @Description  get all users in database
// @Tags         list
// @param 		 keyword   query   string   true   "keyword"
// @param 		 page_size  query  int   true   "pagesize"
// @param 		 page_num   query  int   true   "page num"
// @Accept       json
// @Produce      json
// @Success 0 {object} ListResponse "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
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
			"2022-06-16 20:45:40",},
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
		Status: StatusInfo{
			0,
			"Operation succeed",
		},
		Data:InfoList{
			users,
			1,
			10,
			1,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}
// @Description  get all users in database
// @Tags         list
// @param 		 collection_id   query   string   true   "collection id"
// @Accept       json
// @Produce      json
// @Success 0 {object} ListResponse "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /list/collection [GET]

func SingleColletction(ctx *gin.Context) {
	resp := ListResponse{
		Status: StatusInfo{
			Code: 0,
			Msg: "Operation succeed",
		},
		Data: Collection{
			"mingzheliu-ust-hk",
			"Doodles",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
			"A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Commun",
			[]string{"Comics"},
			20,
			"zezhending-ust-hk",
			"ZZD",
			"2022-06-16 20:45:40",
		},
	}
	ctx.JSON(http.StatusOK, resp)
}
// @Description  get all users in database
// @Tags         list
// @param 		 keyword   path   string   true   "keyword"
// @param 		 rank_favorite   path   string   true   "rank favorite"
// @param 		 rank_time   path   string   true   "rank time"
// @param 		 user_id   path   string   true   "user id"
// @param 		 label   path   string   true   "label"
// @param 		 page_num   path   int   false   "page num"
// @param 		 page_size   path   int   false   "page size"
// @Accept       json
// @Produce      json
// @Success 0 {object} ListResponse "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /list/collection-list [GET]

func CollectionList(ctx *gin.Context) {
	resp := ListResponse{
		Status: StatusInfo{
			0,
			"Operation succeed",
		},
		Data: InfoList{
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

// @Description  get all users in database
// @Tags         list
// @param 		 item_id   query   string   true   "item id"
// @Accept       json
// @Produce      json
// @Success 0 {object} ListResponse "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /list/item [GET]

func SingleItem(ctx *gin.Context) {
	resp := ListResponse{
		Status: StatusInfo{
			0,
			"Operation succeed",
		},
		Data: Item{
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

// @Description  get all users in database
// @Tags         list
// @param 		 page_num   path   int   fasle   "page num"
// @param 		 page_size   path   int   false   "page size"
// @param 		 user_id   path   string   true   "user id"
// @param 		 user_like   path   bool   true   "user like"
// @param 		 user_collect   path   bool   true   "user collect"
// @param 		 user_create   path   bool   true   "user create"
// @param 		 category   path   string   true   "category"
// @param 		 keyword   path   string   true   "keyword"
// @param 		 rank_favorite   path   bool   true   "rank favorite"
// @param 		 rank_time   path   bool   true   "rank time"
// @param 		 collection_id   path   string   true   "collection id"
// @param 		 label   path   string   true   "label"
// @Accept       json
// @Produce      json
// @Success 0 {object} ListResponse "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /list/collection-list [GET]

func ItemList(ctx *gin.Context) {
	resp := ListResponse{
		Status: StatusInfo{
			0,
			"Operation succeed",
		},
		Data: InfoList{
			Info:[]Item{
				{
					"Pixel Bear With Hammer",
					"1010",
					"https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240",
					"2022-06-16 22:04:22",
					"A very cute pixel bear with hammer",
					"Pixel Bear",
					"image",
					[]string{"Music","Comics"},
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
					[]string{"Music","Comics"},
					"mingzheliu-ust-hk",
					"mingzheliu-ust-hk",
					100,
					false,
				},
			} ,
			page: 1,
			size: 10,
			total: 1,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Description  get all users in database
// @Tags         list
// @param 		 item_id   query   string   false   "item id"
// @Accept       json
// @Produce      json
// @Success 0 {object} ListResponse "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /list/collection-list [GET]

func ItemHistory(ctx *gin.Context) {
	resp := ListResponse{
		Status: StatusInfo{
			0,
			"Operation succeed",
		},
		Data: []History{
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


