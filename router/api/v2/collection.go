package v2

import (
	"NFT-BASE-BACK/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCollectionRequest struct {
	Collection_name string `json:"collection_name" example:"Pixel Pear" default:"Pixel Pear"`
	Description     string `json:"description" example:"Happy Happy Happy Happy" default:"Happy Happy Happy Happy"`
	Label           string `json:"label" example:"Pear&Pixel&Wechat" default:"Pear&Pixel&Wechat"`
}
type Collection struct {
	CollectionId   string   `json:"collection_id"`
	CollectionName string   `json:"collection_name"`
	BannerImage    string   `json:"banner_image"`
	LogoImage      string   `json:"logo_image"`
	FeatureImage   string   `json:"feature_image"`
	Description    string   `json:"description"`
	Label          []string `json:"label"`
	ItemNum        int      `json:"item_num"`
	OwnerId        string   `json:"owner_id"`
	OwnerName      string   `json:"owner_name"`
	CreationTime   string   `json:"creation_time"`
}

func NewCollection() Collection {
	return Collection{
		CollectionId:   "111111",
		CollectionName: "Doodles",
		BannerImage:    "https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
		LogoImage:      "https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
		FeatureImage:   "https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a9",
		Description:    "A community-driven collectibles project featuring art by Burnt Toast. Doodles come in a joyful range of colors, traits and sizes with a collection size of 10,000. Each Doodle allows its owner to vote for experiences and activations paid for by the Doodles Community Treasury.",
		Label:          []string{"Comics", "Blue", "Classic", "Scary"},
		ItemNum:        20,
		OwnerId:        "zezhending-ust-hk",
		OwnerName:      "ZZD",
		CreationTime:   "2022-06-16 20:45:40",
	}
}

// Create @Description  create: 创建一个collection
// @Tags         collection
// @param 		 logo_image      formData  file  true    "logo_image of a collection"
// @param 		 feature_image   formData  file  true    "feature_image of a collection"
// @param 		 banner_image    formData  file  true    "banner_image of a collection"
// @param 		 RequestParam  body  CreateCollectionRequest  true  "名称、描述和一些标签，其中标签可以为空，其余参数不可为空"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /collection/create [POST]
func Create(ctx *gin.Context) {
	res := base.Response{}
	res.SetData(NewCollection())
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type EditCollectionRequest struct {
	Collection_id   string `json:"collection_id" example:"111111" default:"111111"`
	Collection_name string `json:"collection_name" example:"Pixel Pear" default:"Pixel Pear"`
	Description     string `json:"description" example:"Happy Happy Happy Happy" default:"Happy Happy Happy Happy"`
	Label           string `json:"label" example:"Pear&Pixel&Wechat" default:"Pear&Pixel&Wechat"`
}

// Edit  @Description  edit : 输入邮箱、验证码和密码，重新设置已经忘记的密码
// @Tags         collection
// @param 		 logo_image      formData  file  false    "logo_image of a collection"
// @param 		 feature_image   formData  file  false    "feature_image of a collection"
// @param 		 banner_image    formData  file  false    "banner_image of a collection"
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
	res.SetData(NewCollection())
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}
