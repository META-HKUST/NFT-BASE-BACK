package v1

import (
	"NFT-BASE-BACK/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type JsonMsg struct {
	Name  string `josn:"name" example:"ust-hk #9897"`
	Image string `json:"image" exmaple:"https://ikzttp.mypinata.cloud/ipfs/QmYDvPAXtiJg7s8JdRBSLWdgSphQdac8j1YuQNNxcGE1hg/9897.png"`
}

// GetJsonMsg @Description  item-json: get one item json info related to ipfs
// @Tags         item
// @param 		 token-id   query   string    true    "token id of one NFT"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /items/item-json [GET]
func GetJsonMsg(ctx *gin.Context) {
	id := ctx.Query("token-id")
	tokenId, err := strconv.ParseInt(id, 10, 64)
	tokenInfo, err := model.GetUrlByTokenId(tokenId)
	if err != nil {
		return
	}
	fmt.Println(tokenInfo)
	resp := JsonMsg{
		Name:  "ust-hk #9897",
		Image: "https://ikzttp.mypinata.cloud/ipfs/QmYDvPAXtiJg7s8JdRBSLWdgSphQdac8j1YuQNNxcGE1hg/9897.png",
	}
	ctx.JSON(http.StatusOK, resp)

	model.StoreUrl(4, "http")
}
