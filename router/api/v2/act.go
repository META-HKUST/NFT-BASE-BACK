package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostActCreateRequest struct {
	ActName     string `json:"act_name" example:"the first activity"`
	Description string `json:"description" example:"It is funny"`
	StartTime   string `json:"start_time" example:"2022-06-18 20:45:40"`
	EndTime     string `json:"end_time" example:"2022-06-20 20:45:40"`
}

// PostActCreate
// @Description  create activity
// @Tags         act
// @param        RequestParam body     PostActCreateRequest true "参数均不可为空"
// @param 		 act_image   formData  file  true   "activity front page image"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /act/create [POST]
func PostActCreate(ctx *gin.Context) {
	// var req PostActCreateRequest
	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, err.Error())
	// 	return
	// }
	// log.Println(req)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Operation succeed",
		"data": gin.H{
			"act_id":      1,
			"act_name":    "FIRST ACT",
			"creater_id":  "zwang-ust-hk",
			"create_time": "2022-06-18 20:45:40",
			"start_time":  "2022-06-20 20:45:40",
			"end_time":    "2022-08-20 20:45:40",
			"act_image":   "https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F5c-HcdLMinTg3LvEwXYZYC-u5nN22Pn5ivTPYA4pVEsWJHU1rCobhUlHSFjZgCHPGSmcGMQGCrDCQU8BfSfygmL7Uol9MRQZt6-gqA%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=b61af932ff80dfea857abd3a4650f4f2",
			"description": "FIRST ACT",
			"item_num":    10,
		},
	})
}

type PostActDeleteRequest struct {
	ActId string `json:"act_id" example:"1"`
}

// PostActDelete
// @Description  deleta activity
// @Tags         act
// @param        RequestParam body PostActDeleteRequest true "活动ID不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/delete [POST]
// @Security ApiKeyAuth
func PostActDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Operation succeed",
		"data": gin.H{},
	})
}

type PostActEditRequest struct {
	ActID       string `json:"act_id" example:"1"`
	ActName     string `json:"act_name" example:"the first activity"`
	Description string `json:"description" example:"It is funny"`
	StartTime   string `json:"start_time" example:"2022-06-20 20:45:40"`
	EndTime     string `json:"end_time" example:"2022-08-20 20:45:40"`
}

// PostActEditRequest
// @Description  edit activity
// @Tags         act
// @param        RequestParam body     PostActEditRequest true "ID不可为空，其他可为空"
// @param 		 act_image   formData  file  false   "activity front page image"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/edit [POST]
// @Security ApiKeyAuth
func PostActEdit(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Operation succeed",
		"data": gin.H{
			"act_id":      1,
			"act_name":    "FIRST ACT",
			"creater_id":  "zwang-ust-hk",
			"create_time": "2022-06-18 20:45:40",
			"start_time":  "2022-06-20 20:45:40",
			"end_time":    "2022-08-20 20:45:40",
			"act_image":   "https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F5c-HcdLMinTg3LvEwXYZYC-u5nN22Pn5ivTPYA4pVEsWJHU1rCobhUlHSFjZgCHPGSmcGMQGCrDCQU8BfSfygmL7Uol9MRQZt6-gqA%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=b61af932ff80dfea857abd3a4650f4f2",
			"description": "FIRST ACT",
			"item_num":    10,
		},
	})
}

// GetActInfo
// @Description  get activity information
// @Tags         act
// @param        act_id 	query	string	true 	"ID不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/info [GET]
func GetActInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Operation succeed",
		"data": gin.H{
			"act_id":      1,
			"act_name":    "FIRST ACT",
			"creater_id":  "zwang-ust-hk",
			"create_time": "2022-06-18 20:45:40",
			"start_time":  "2022-06-20 20:45:40",
			"end_time":    "2022-08-20 20:45:40",
			"act_image":   "https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F5c-HcdLMinTg3LvEwXYZYC-u5nN22Pn5ivTPYA4pVEsWJHU1rCobhUlHSFjZgCHPGSmcGMQGCrDCQU8BfSfygmL7Uol9MRQZt6-gqA%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=b61af932ff80dfea857abd3a4650f4f2",
			"description": "FIRST ACT",
			"item_num":    10,
		},
	})
}

type PostActUploadItemRequest struct {
	ActID  string `json:"act_id" example:"1"`
	ItemID string `json:"item_id" example:"2"`
}

// PostActUploadItem
// @Description  upload item to activity, means students join this activity
// @Tags         act
// @param        RequestParam body	PostActUploadItemRequest true "ID均不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/upload-item [POST]
// @Security ApiKeyAuth
func PostActUploadItem(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Operation succeed",
		"data": gin.H{},
	})
}

// GetActItemList
// @Description  get activity item list
// @Tags         act
// @param        act_id 	query	string 		true 	"act id"
// @param        page_num 	query 	int 		true 	"page num"
// @param        page_size 	query 	int 		true 	"page size"
// @param        rank_vote 	query 	boolean 	true 	"whether sorted by votes or not"
// @param        rang_time 	query 	boolean 	true 	"whether sorted by time or not"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/item-list [GET]
func GetActItemList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Operation succeed",
		"data": gin.H{
			"item_list": []gin.H{
				{
					"item_name":     "Pixel Bear With Hammer",
					"item_id":       "1010",
					"item_data":     "https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240",
					"create_time":   "2022-06-16 22:04:22",
					"description":   "A very cute pixel bear with hammer",
					"collection_id": "Pixel Bear",
					"category":      "image",
					"label": []string{
						"Music",
						"Comics",
					},
					"creater_id": "mingzheliu-ust-hk",
					"owner_id":   "mingzheliu-ust-hk",
					"vote_num":   100,
					"vote":       false,
				},
				{
					"item_name":     "Pixel Bear With Hammer",
					"item_id":       "1010",
					"item_data":     "https://img1.baidu.com/it/u=1783064339,1648739044&fm=253&fmt=auto&app=138&f=GIF?w=240&h=240",
					"create_time":   "2022-06-16 22:04:22",
					"description":   "A very cute pixel bear with hammer",
					"collection_id": "Pixel Bear",
					"category":      "image",
					"label": []string{
						"Music",
						"Comics",
					},
					"creater_id": "mingzheliu-ust-hk",
					"owner_id":   "mingzheliu-ust-hk",
					"vote_num":   100,
					"vote":       false,
				},
			},
			"page":  1,
			"size":  10,
			"total": 1,
		},
	})
}

type PostActVoteRequest struct {
	ActID   string `json:"act_id" example:"1"`
	ItemID  string `json:"item_id" example:"2"`
	VoteNum int    `json:"vote_num" example:"3" default:"1"`
}

// PostActVote
// @Description  student votes for item in activity
// @Tags         act
// @param        RequestParam body	PostActVoteRequest true "参数均不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/vote [POST]
// @Security ApiKeyAuth
func PostActVote(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Operation succeed",
		"data": gin.H{},
	})
}
