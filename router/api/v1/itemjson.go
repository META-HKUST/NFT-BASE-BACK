package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonMsg struct {
	Name		string `josn:"name" example:"ust-hk #9897"`
	Image       string `json:"image" exmaple:"https://ikzttp.mypinata.cloud/ipfs/QmYDvPAXtiJg7s8JdRBSLWdgSphQdac8j1YuQNNxcGE1hg/9897.png"`
}
// @Description  item json msg
// @Tags         item
// @param 		 item-id   path   string    true    "collection id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  Item       "GET/api/v1/XXXX"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /{item-id} [GET]
func GetJsonMsg(ctx *gin.Context)  {
	resp := JsonMsg{
		Name:	"ust-hk #9897",
		Image:  "https://ikzttp.mypinata.cloud/ipfs/QmYDvPAXtiJg7s8JdRBSLWdgSphQdac8j1YuQNNxcGE1hg/9897.png",
	}
	ctx.JSON(http.StatusOK, resp)
}

