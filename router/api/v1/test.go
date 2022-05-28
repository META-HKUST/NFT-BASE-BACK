package v1

import (
	"NFT-BASE-BACK/sdk"
	"net/http"

	"github.com/gin-gonic/gin"
)

type submitReq struct {
	Username     string   `json:"username"`
	ContractName string   `json:"contract"`
	Args         []string `json:"args"`
}

func TestContract(ctx *gin.Context) {
	var req submitReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 9998,
			"msg":  err,
		})
		return
	}
	result, err := sdk.Submit(req.Username, req.ContractName, req.Args...)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 9999,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}

type enrollReq struct {
	Username string `json:"username"`
}

func TestEnroll(ctx *gin.Context) {
	var req enrollReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 9998,
			"msg":  err,
		})
		return
	}

	result, err := sdk.Enroll(req.Username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 9999,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
