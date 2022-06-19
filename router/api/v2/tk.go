package v2

import (
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
	// var req PostTokenTransferRequest
	// err := ctx.ShouldBindJSON(&req)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, err.Error())
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Operation succeed",
		"data": gin.H{
			"user_id": "mazhengwang-ust-hk",
			"token":   100,
		},
	})
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
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Operation succeed",
		"data": gin.H{
			"user_id": "mazhengwang-ust-hk",
			"token":   100,
		},
	})
}
