package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMetaData
// @Description  get meta data information
// @Tags         metadata
// @param 		 token_id   query   string   true   "token_id"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /metadata [GET]
func GetMetaInfo(ctx *gin.Context) {
	var SuccessResp base.TokenUrlResponse
	token_id := ctx.Query("token_id")

	metaInfo,err := service.GetMetaInfo(token_id)
	if err != nil {
		ctx.JSON(http.StatusOK, base.ServerError)
		return
	}
	SuccessResp.Image = metaInfo.IpfsUrl
	SuccessResp.Name = metaInfo.Name
	ctx.JSON(http.StatusOK, SuccessResp)
}