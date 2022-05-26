package v1

import (
	"NFT-BASE-BACK/sdk"
	"net/http"

	"github.com/gin-gonic/gin"
)

type submitReq struct {
	Username     string   `json:"username" binding:"min"`
	ContractName string   `json:"contract"`
	Args         []string `json:"args"`
}

func TestContract(ctx *gin.Context) {
	var req submitReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 9999,
			"msg":  err,
		})
		return
	}
	result, err := sdk.Submit(req.Username, req.ContractName, req.Args...)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 9999,
			"msg":  err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}
