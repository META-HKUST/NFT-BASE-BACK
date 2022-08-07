package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/entity"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/service"
	"NFT-BASE-BACK/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type CreateCollectionRequest struct {
	Collection_name string   `json:"collection_name" example:"Pixel Pear" default:"Pixel Pear"`
	LogoImage       string   `json:"logo_image" default:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	LogoImage_S     string   `json:"log_image_signature" default:"ABCDE"`
	FeatureImage    string   `json:"feature_image" default:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	FeatureImagee_S string   `json:"feature_image_signature" default:"ABCDE"`
	BannerImage     string   `json:"banner_image" default:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9P"`
	BannerImage_S   string   `json:"banner_image_signature" default:"ABCDE"`
	Description     string   `json:"description" example:"Happy Happy Happy Happy" default:"Happy Happy Happy Happy"`
	Label           []string `json:"label" example:"Music,Comics"`
}

// Create @Description  create: 创建一个collection
// @Tags         collection
// @param 		 RequestParam  body  CreateCollectionRequest  true  "名称、描述和一些标签，其中标签可以为空，其余参数不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /collection/create [POST]
func Create(ctx *gin.Context) {
	ch := CreateCollectionRequest{}
	ctx.BindJSON(&ch)
	res := base.Response{}

	// check empty
	ss := append([]string{}, ch.Collection_name, ch.LogoImage, ch.FeatureImage, ch.BannerImage, ch.Description)
	if utils.CheckAnyEmpty(ss) == false {
		ctx.JSON(http.StatusOK, res.SetCode(base.EmptyInput))
		return
	}

	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)

	code, data := service.CreateCollectionByAccount(UserId, ch.Collection_name, ch.LogoImage, ch.FeatureImage, ch.BannerImage, ch.Description)

	// 写item label
	for _, v := range ch.Label {
		CollectionLabel := model.CollectionLabel{
			// change this to test
			CollectionID:    data.CollectionId,
			CollectionLabel: v,
		}
		_, err := model.CreateCollectionLabel(CollectionLabel)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, "database error")
			return
		}
	}

	res.SetData(data)
	ctx.JSON(http.StatusOK, res.SetCode(code))
}

type EditCollectionRequest struct {
	Collection_id   int      `json:"collection_id" example:"111111" default:"111111"`
	Collection_name string   `json:"collection_name" example:"Pixel Pear" default:"Pixel Pear"`
	LogoImage       string   `json:"logo_image" default:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	LogoImage_S     string   `json:"log_image_signature" default:"ABCDE"`
	FeatureImage    string   `json:"feature_image" default:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9"`
	FeatureImagee_S string   `json:"feature_image_signature" default:"ABCDE"`
	BannerImage     string   `json:"banner_image" default:"https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9P"`
	BannerImage_S   string   `json:"banner_image_signature" default:"ABCDE"`
	Description     string   `json:"description" example:"Happy Happy Happy Happy" default:"Happy Happy Happy Happy"`
	Label           []string `json:"label" example:"Music,Comics"`
}

// Edit  @Description  edit : 输入邮箱、验证码和密码，重新设置已经忘记的密码
// @Tags         collection
// @param 		 RequestParam  body  EditCollectionRequest  true  "输入想要修改的名称、描述和标签 collection_id不能为空，其他都可以为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /collection/edit [POST]
func Edit(ctx *gin.Context) {
	res := base.Response{}

	ch := EditCollectionRequest{}
	ctx.BindJSON(&ch)
	// check empty
	if utils.CheckIntEmpty(ch.Collection_id) == false {
		resp := base.Response{}
		ctx.JSON(http.StatusOK, resp.SetCode(base.EmptyInput))
		return
	}
	s, _ := ctx.Get("email")

	email := fmt.Sprintf("%v", s)

	t1 := strings.Replace(email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)

	// 检查collection是否是这个人拥有的
	col, _ := model.GetCollection(ch.Collection_id)
	if col.Owner != UserId {
		ctx.JSON(http.StatusOK, res.SetCode(base.PermissionDenied))
		return
	}

	log.Println("Edit Collection UserId: ", UserId)

	c := entity.Collection{
		CollectionId:   ch.Collection_id,
		CollectionName: ch.Collection_name,
		LogoImage:      ch.LogoImage,
		FeatureImage:   ch.FeatureImage,
		BannerImage:    ch.BannerImage,
		Description:    ch.Description,
	}

	code, data := service.EditCollection(c)

	if ch.Label != nil {
		err := model.EditCollectionLable(ch.Label, ch.Collection_id)
		if err != nil {
			ctx.JSON(http.StatusOK, base.ServerError)
		}
	}

	res.SetData(data)

	ctx.JSON(http.StatusOK, res.SetCode(code))
}
