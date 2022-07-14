package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/service"
	"NFT-BASE-BACK/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostTokenTransferRequest struct {
	TokenNum   float32 `json:"token_num" example:"100" binding:"required"`
	FromUserId string  `json:"from_user_id" example:"mazhengwang-ust-hk" binding:"required"`
	ToUserId   string  `json:"to_user_id" example:"mzliu" binding:"required"`
}

// PostTokenTransfer
// @Description  transfer token form one account to another account
// @Tags         tk
// @param 		 param_request   body   PostTokenTransferRequest   true   "参数均不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /tk/transfer [POST]
// @Security ApiKeyAuth
func PostTokenTransfer(ctx *gin.Context) {
	req := PostTokenTransferRequest{}
	res := base.Response{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// check empty
	ss := append([]string{}, req.FromUserId, req.ToUserId)
	if utils.CheckAnyEmpty(ss) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}
	// check empty
	if req.TokenNum == 0 {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}

	tokenInfo, errCode := service.Transfer(req.TokenNum, req.FromUserId, req.ToUserId)
	if errCode != base.Success {
		ctx.JSON(http.StatusOK, res.SetCode(errCode))
		return
	}
	res.SetData(tokenInfo)
	ctx.JSON(http.StatusOK, res.SetCode(errCode))
}

// GetTokenInfo
// @Description get token balance
// @Tags         tk
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /tk/info [GET]
// @Security ApiKeyAuth
func GetTokenInfo(ctx *gin.Context) {
	res := base.Response{}
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, "auth email error")
		return
	}

	tokenInfo, code := service.GetTokenInfo(email.(string))
	if code != base.Success {
		ctx.JSON(http.StatusOK, res.SetCode(code))
		return
	}

	res.SetData(tokenInfo)
	ctx.JSON(http.StatusOK, res.SetCode(code))
}
