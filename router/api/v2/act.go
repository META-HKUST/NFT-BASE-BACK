package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PostActCreateRequest struct {
	ActName     string `json:"act_name" example:"the first activity" `
	Description string `json:"description" example:"It is funny"`
	StartTime   string `json:"start_time" example:"2022-06-18 20:45:40"`
	EndTime     string `json:"end_time" example:"2022-06-20 20:45:40" `
	ActImage    string `json:"act_image" example:"abc.com" `
	ActImage_S  string `json:"act_image_signature" example:"abc.com" `
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

	//if utils.CheckEmptyStruct(ch) == false {
	//	res.SetCode(base.EmptyInput)
	//	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
	//	return
	//}

	ss := append([]string{}, ch.ActName, ch.StartTime, ch.EndTime, ch.ActImage, ch.Description)
	if utils.CheckAnyEmpty(ss) == false {
		ctx.JSON(http.StatusOK, res.SetCode(base.EmptyInput))
		return
	}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	// check if admin account
	if email != "admin@unifit.art" {
		ctx.JSON(http.StatusOK, base.PermissionDenied)
	}

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

	if utils.CheckIntEmpty(ch.ActId) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}

	res := base.Response{}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	// check if admin account
	if email != "admin@unifit.art" {
		ctx.JSON(http.StatusOK, base.PermissionDenied)
	}

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
	if utils.CheckIntEmpty(ch.ActId) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}

	res := base.Response{}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	// check if admin account
	if email != "admin@unifit.art" {
		ctx.JSON(http.StatusOK, base.PermissionDenied)
	}

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
	if utils.CheckEmpty(actId) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}
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

	if utils.CheckIntEmpty(ch.ActID) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}
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

	err := model.UploadNFT(ch.ActID, ch.ItemID)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
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

	// check empty
	if utils.CheckIntEmpty(ch.ActID) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}
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

	var res base.Response

	data, err := model.GetAllAct()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(data)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

// GetActItemList
// @Description  get activity item list
// @Tags         act
// @param        act_id 	query	string 		true 	"act id"
// @param        page_num 	query 	int 		true 	"page num"
// @param        page_size 	query 	int 		true 	"page size"
// @param        keyword 	query 	string 		false 	"keyword"
// @param        rank_vote 	query 	boolean 	false 	"whether sorted by votes or not"
// @param        rank_time 	query 	boolean 	false 	"whether sorted by time or not"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/item-list [GET]
func GetActItemList(ctx *gin.Context) {

	var res base.Response

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
	rankTime := ctx.Query("rank_favorite")
	rankTimeBool, _ := strconv.ParseBool(rankTime)
	rankVote := ctx.Query("rank_vote")
	rankVoteBool, _ := strconv.ParseBool(rankVote)
	actId := ctx.Query("act_id")
	actIdInt, _ := strconv.ParseInt(actId, 10, 64)
	keyword := ctx.Query("keyword")

	data, err := model.GetActItemList(pageNumInt, pageSizeInt, actIdInt, rankVoteBool, rankTimeBool, UserId, keyword)

	if err != nil {
		log.Println(err)
		res.SetData(data)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(data)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

// CanUpload
// @Description  get the items that this user could upload to this act
// @Tags         act
// @param        page_num 	query 	int 		true 	"page num"
// @param        page_size 	query 	int 		true 	"page size"
// @param        act_id 	query	string 		true 	"act id"
// @param        user_id 	query 	string		true 	"user id"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /act/can-upload [GET]
func CanUpload(ctx *gin.Context) {

	var res base.Response

	actId := ctx.Query("act_id")
	actIdInt, _ := strconv.ParseInt(actId, 10, 64)
	userId := ctx.Query("user_id")
	pageNum := ctx.Query("page_num")
	pageNumInt, _ := strconv.ParseInt(pageNum, 10, 64)
	pageSize := ctx.Query("page_size")
	pageSizeInt, _ := strconv.ParseInt(pageSize, 10, 64)

	data, err := model.GetCanUpload(pageNumInt, pageSizeInt, actIdInt, userId)
	if err != nil {
		log.Println(err)
		res.SetData(data)
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(data)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}
