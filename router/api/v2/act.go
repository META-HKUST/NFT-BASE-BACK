package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PostActCreateRequest struct {
	ActName     string `json:"act_name" example:"the first activity"`
	Description string `json:"description" example:"It is funny"`
	StartTime   string `json:"start_time" example:"2022-06-18 20:45:40"`
	EndTime     string `json:"end_time" example:"2022-06-20 20:45:40"`
	ActImage    string `json:"act_image" example:"abc.com"`
}

// PostActCreate
// @Description  create activity
// @Tags         act
// @param        RequestParam body     PostActCreateRequest true "参数均不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /act/create [POST]
func PostActCreate(ctx *gin.Context) {
	ch := PostActCreateRequest{}
	ctx.BindJSON(&ch)

	res := base.Response{}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)

	data, err := model.AddAction(ch.ActName, UserId, ch.StartTime, ch.EndTime, ch.ActImage, ch.Description, 0)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(data)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type ActRequest struct {
	ActId int `json:"act_id" example:"1"`
}

// PostActDelete
// @Description  delet activity
// @Tags         act
// @param        RequestParam body ActRequest true "活动ID不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/delete [POST]
// @Security ApiKeyAuth
func PostActDelete(ctx *gin.Context) {
	ch := ActRequest{}
	ctx.BindJSON(&ch)

	res := base.Response{}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)
	fmt.Println(UserId)

	err := model.DeleteAction(ch.ActId)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type PostActEditRequest struct {
	ActId       int    `json:"act_id" example:"1"`
	ActName     string `json:"act_name" example:"the first activity"`
	Description string `json:"description" example:"It is funny"`
	StartTime   string `json:"start_time" example:"2022-06-18 20:45:40"`
	EndTime     string `json:"end_time" example:"2022-06-20 20:45:40"`
	ActImage    string `json:"act_image" example:"abc.com"`
}

// PostActEdit
// @Description  edit activity
// @Tags         act
// @param        RequestParam body     PostActEditRequest true "ID不可为空，其他可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/edit [POST]
// @Security ApiKeyAuth
func PostActEdit(ctx *gin.Context) {
	ch := PostActEditRequest{}
	ctx.BindJSON(&ch)

	res := base.Response{}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)
	fmt.Println(UserId)

	data, err := model.EditAction(ch.ActId, ch.ActName, ch.StartTime, ch.EndTime, ch.ActImage, ch.Description)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(data)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

// GetActInfo
// @Description  get activity information
// @Tags         act
// @param        act_id 	query	int 	true 	"ID不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/info [GET]
func GetActInfo(ctx *gin.Context) {
	actId := ctx.Query("act_id")

	res := base.Response{}

	data, err := model.GetAction(actId)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(data)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type PostActUploadItemRequest struct {
	ActID  int    `json:"act_id" example:"1"`
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
	ch := PostActUploadItemRequest{}
	ctx.BindJSON(&ch)

	res := base.Response{}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)
	fmt.Println(UserId)

	err := model.UploadNFT(ch.ActID, ch.ItemID)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

// GetActItemList
// @Description  get activity item list
// @Tags         act
// @param        act_id 	query	string 		true 	"act id"
// @param        page_num 	query 	int 		true 	"page num"
// @param        page_size 	query 	int 		true 	"page size"
// @param        rank_vote 	query 	boolean 	true 	"whether sorted by votes or not"
// @param        rank_time 	query 	boolean 	true 	"whether sorted by time or not"
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
	ActID  int    `json:"act_id" example:"1"`
	ItemID string `json:"item_id" example:"2"`
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
	ch := PostActUploadItemRequest{}
	ctx.BindJSON(&ch)

	res := base.Response{}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)
	fmt.Println(UserId)

	bo, _ := model.DoesVote(ch.ActID, ch.ItemID, UserId)
	if bo == false {
		err := model.Vote(ch.ActID, ch.ItemID, UserId)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
			return
		}
		ctx.JSON(http.StatusOK, res.SetCode(base.Success))
	} else {
		err := model.UnVote(ch.ActID, ch.ItemID, UserId)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
			return
		}
		ctx.JSON(http.StatusOK, res.SetCode(base.Success))
	}
}

type ActRes struct {
	ActId       int    `json:"act_id" example:"1"`
	ActName     string `json:"act_name" example:"the first activity"`
	CreaterId   string `json:"creater_id" example:"ssscaxxadw05130"`
	CreateTime  string `json:"create_time" example:"2022-06-18 20:45:40"`
	StartTime   string `json:"start_time" example:"2022-06-18 20:45:40"`
	EndTime     string `json:"end_time" example:"2022-06-20 20:45:40"`
	ActImage    string `json:"act_image" example:"abc.com"`
	Description string `json:"description" example:"It is funny"`
	ItemNum     int    `json:"item_num" example:"100"`
}

type ActCount struct {
	ActCount int `json:"act_count" example:"100"`
}

// GetActCount
// @Description  get action count
// @Tags         act
// @Accept       json
// @Produce      json
// @Success 200  {object}   ActCount  "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/act-count [GET]
func GetActCount(ctx *gin.Context) {

	res := base.Response{}

	data, err := model.GetActionCount()

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(data)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

// GetAllAct
// @Description  get all act
// @Tags         act
// @Accept       json
// @Produce      json
// @Success 200  {object}   ActRes  "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/all-action [GET]
func GetAllAct(ctx *gin.Context) {

	res := base.Response{}

	data, err := model.GetAllAct()

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(data)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}
