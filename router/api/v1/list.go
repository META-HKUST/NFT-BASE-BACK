package v1

import (
	"NFT-BASE-BACK/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListResponse struct {
	Code			int
	Msg 			string
	Data 			ListResponseBody
}

type UserInfo struct {
	UserId				string		`json:"user_id" example:"zhengwang-ust-hk"`
	UserEmail			string		`json:"user_email" example:"zhengwang@ust.hk"`
	UserName			string		`json:"user_name"`
	BannerImage			string		`json:"banner_image"`
	LogoImage			string		`json:"logo_image"`
	Poison				string		`json:"poison"`
	Organization		string		`json:"organization"`
	RegisterationTime	string		`json:"registeration_time"`
}

type CollectionIfo struct {
	CollectionId		string		`json:"collection_id"`
	CollectionName		string		`json:"collection_name"`
	BannerImage			string		`json:"banner_image"`
	LogoImage			string		`json:"logo_image"`
	FeatureImage		string		`json:"feature_image"`
	Description			string		`json:"description"`
	Label				[]string		`json:"label"`
	ItemNum				int64		`json:"item_num"`
	OwnerId				string		`json:"owner_id"`
	OwnerName			string		`json:"owner_name"`
	Creation_time		string		`json:"creation_time"`
}

type Collection struct {
	CollectionName		string		`json:"collection_name"`
	BannerImage			string		`json:"banner_image"`
	LogoImage			string		`json:"logo_image"`
	FeatureImage		string		`json:"feature_image"`
	Description			string		`json:"description"`
	Label				[]string		`json:"label"`
	ItemNum				int64		`json:"item_num"`
	OwnerId				string		`json:"owner_id"`
	OwnerName			string		`json:"owner_name"`
	Creation_time		string		`json:"creation_time"`
}


type History struct {
	FromUserId		string
	ToUserId		string
	FromUserName	string
	ToUserName		string
	Time			string
}

type ListResponseBody struct {
	Info 	[]UserInfo
	Page 	int64
	Size	int64
	Total 	int64
}


type CollectListBody struct {
	Info 	[]Collection
	Page 	int64
	Size	int64
	Total 	int64
}

type ItemListBody struct {
	Item 	[]NewItem
	Page 	int64
	Size	int64
	Total 	int64
}


// @Description
// @Tags         list
// @param 		 keyword   query   string    true    "keyword"
// @param 		 page_size  query  int    true    "page size"
// @param 		 page_num   query   int    true    "page num"
// @Accept       json
// @Produce      json
// @Success 0 {object} ListResponse "Operation Succeed, code: 0"
// @Failure 400 {object} string "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /list/user-list [GET]

func UserList(ctx *gin.Context) {
	//Keyword := ctx.Query("keyword")

	PageSize, err := strconv.ParseInt(ctx.Query("page_size"), 10, 64)
	PageNum,err :=  strconv.ParseInt(ctx.Query("page_num"), 10, 64)
	if err != nil {
		return
	}

	ItemList := []UserInfo{
		{
			"mingzheliu-ust-hk",
			"mingzheliu@ust.hk",
			"LMZ",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
			"teacher",
			"HKUST-GZ",
			utils.GetTimeNow(),
		},
		{
			"zezhending-ust-hk",
			"zezhending@ust.hk",
			"ZZD",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
			"student",
			"HKUST",
			"2022-06-16 20:45:40",
		},
	}
	resbody := ListResponseBody{
	ItemList,
	PageNum,
	PageSize,
	1,
	}

	resp := ListResponse{
		0,
		"Operation succeed",
		resbody,
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Description
// @Tags         list
// @param 		 collection_id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /list/collection [GET]
func SingleCollection(ctx *gin.Context) {

	collectionInfo := CollectionIfo{
		"111111",
		"Doodles",
		"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
		"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
		"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
		"A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Community Treasury.",
		[]string{"Comics"},
		20,
		"zezhending-ust-hk",
		"ZZD",
		utils.GetTimeNow(),
	}

	resp := struct{
		Code 	int
		Msg 	string
		Data 	CollectionIfo
	}{
		0,
		"Operation succeed",
		collectionInfo,

	}
	ctx.JSON(http.StatusOK, resp)
}

// @Description
// @Tags         list
// @param 		 collection_id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /list/collection-list [GET]

func CollectionList(ctx *gin.Context) {

	CollectionList := []Collection{
		{"Doodles",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
			"A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Community Treasury.",
			[]string{"Comics"},
			20,
			"zezhending-ust-hk",
			"ZZD",
			"2022-06-16 20:45:40",
		},
		{
			"MAYC",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F5c-HcdLMinTg3LvEwXYZYC-u5nN22Pn5ivTPYA4pVEsWJHU1rCobhUlHSFjZgCHPGSmcGMQGCrDCQU8BfSfygmL7Uol9MRQZt6-gqA%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=b61af932ff80dfea857abd3a4650f4f2",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2FlHexKRMpw-aoSyB1WdFBff5yfANLReFxHzt1DOj_sg7mS14yARpuvYcUtsyyx-Nkpk6WTcUPFoG53VnLJezYi8hAs0OxNZwlw6Y-dmI%3Ds10000?fit=max&h=120&w=120&auto=format&s=d21114ca201b6479e28180b672436109",
			"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F5c-HcdLMinTg3LvEwXYZYC-u5nN22Pn5ivTPYA4pVEsWJHU1rCobhUlHSFjZgCHPGSmcGMQGCrDCQU8BfSfygmL7Uol9MRQZt6-gqA%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=b61af932ff80dfea857abd3a4650f4f2",
			"The MUTANT APE YACHT CLUB is a collection of up to 20,000 Mutant Apes that can only be created by exposing an existing Bored Ape to a vial of MUTANT SERUM or by minting a Mutant Ape in the public sale.",
			[]string{"Comics"},
			20,
			"zwang-ust-hk",
			"ZW",
			"1655981865",
		},
	}

	resp := struct {
		Code 	int
		Msg		string
		Data 	CollectListBody
	}{
		0,
		"Operation succeed",
		CollectListBody{
			CollectionList,
			1,
			10,
			1,
		},
	}
	ctx.JSON(http.StatusOK, resp)


}
// @Description
// @Tags         list
// @param 		 item_id   path   string    true    "item id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /list/item [GET]
func SingleItem(ctx *gin.Context) {
	ItemID := ctx.Query("item_id")

	ItemInfo := NewItem{
		"Pixel Bear With Hammer",
		ItemID,
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
	}

	resp := struct {
		Code int
		Msg string
		Data NewItem
	}{
		0,
		"Operation succeed",
		ItemInfo,
	}
	ctx.JSON(http.StatusOK, resp)
}
// @Description
// @Tags         list
// @param 		 page_num   path   int    false    "page num"
// @param 		 page_size   path   int    false    "page size"
// @param 		 user_id   path   string    true    "user id"
// @param 		 user_like   path   bool    true    "user like"
// @param 		 user_collect   path   bool    true    "user collect"
// @param 		 user_create   path   bool    true    "user create"
// @param 		 category   path   string    true    "category"
// @param 		 keyword   path   string    true    "keyword"
// @param 		 rank_favorite   path   bool    true    "rank favorite"
// @param 		 rank_time   path   bool    true    "rank time"
// @param 		 collection_id   path   string    true    "collection id"
// @param 		 label   path   string    true    "label"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /list/item-list [GET]
func ItemList(ctx *gin.Context) {

	ItemList := []NewItem{
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
	}

	resp := struct {
		Code int
		Msg string
		Data ItemListBody
	}{
		0,
		"Operation succeed",
		ItemListBody{
			ItemList,
			1,
			10,
			1,
		},
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Description
// @Tags         list
// @param 		 item_id   path   string    true    "item id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /list/item_history [GET]
func ItemHistory(ctx *gin.Context) {
	//ItemID := ctx.Query("item_id")

	HistoryList := []History{
		{
			"baofuhan-ust-hk",
			"zwang-ust-hk",
			"BFH",
			"ZW",
			"2022-06-16 20:45:40",
		},{
			"baofuhan-ust-hk",
			"zwang-ust-hk",
			"BFH",
			"ZW",
			"2022-06-16 20:45:40",
		},
	}

	resp := struct {
		Code int
		Msg string
		Data []History
	}{
		0,
		"Operation succeed",
		HistoryList,
	}
	ctx.JSON(http.StatusOK, resp)
}
