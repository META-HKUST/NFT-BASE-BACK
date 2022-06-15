package v1

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateItem @Description  create single item: parse UserId from token and create NFT(Creater and Owner are defined as UserId)
// @Tags         account
// @param 		 name   query   string   true   "name"
// @param 		 image   query   string   true    "image"
// @param 		 description   query   string   true    "description"
// @param 		 collection   query   string   true    "collection"
// @param 		 category   query   string   true    "category"
// @param 		 label   query   string   true    "label"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /account/create-item [POST]
// @Security ApiKeyAuth
func CreateItem(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	name := ctx.Query("name")
	image := ctx.Query("image")
	description := ctx.Query("description")
	itemCollection := ctx.Query("collection")
	category := ctx.Query("category")
	label := ctx.Query("label")
	code, data := service.CreateItem(UserId, name, image, description, itemCollection, category, []string{label})
	res.SetCode(code)
	res.SetData(data)
	ctx.JSON(http.StatusOK, res)
}

// EditItem @Description  edit single item: parse UserId from token and edit NFT
// @Tags         account
// @param 		 name   query   string   false   "name"
// @param 		 description   query   string   false   "description"
// @param 		 collection   query   string   false   "collection"
// @param 		 label   query   string   false   "label"
// @param 		 item-id   query   string   false   "item-id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /account/edit-item [POST]
// @Security ApiKeyAuth
func EditItem(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	name := ctx.Query("name")
	item_id := ctx.Query("item-id")
	description := ctx.Query("description")
	itemCollection := ctx.Query("collection")
	label := ctx.Query("label")
	image := "/yezzi/1.png"
	category := "image"
	fmt.Println(item_id)
	code, data := service.EditItem(UserId, name, image, description, itemCollection, category, []string{label})
	res.SetCode(code)
	res.SetData(data)
	ctx.JSON(http.StatusOK, res)
}

// Collected @Description  collected: will return the collected NFT items of one user, enter pagesize and pagenumber return number and info
// @Tags         account
// @param 		 num   query   string   true   "num"
// @param 		 pagesize   query   string   true   "pagesize"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /account/collected [GET]
// @Security ApiKeyAuth
func Collected(ctx *gin.Context) {
	res := base.PageResponse{}
	UserId := "mingzheliu-ust-hk"
	pgnumber := ctx.Query("num")
	pgsize := ctx.Query("pagesize")
	fmt.Println(pgnumber, pgsize)
	code, data, count := service.GetItems(UserId, 1, 2, "collected")
	res.SetCode(code)
	res.SetData(data)
	res.SetCount(count)
	ctx.JSON(http.StatusOK, res)
}

// Favorites @Description  favorites: will return the favorite NFT items of one user, in which UserId is parsed from token
// @Tags         account
// @param 		 num   query   string   true   "num"
// @param 		 pagesize   query   string   true   "pagesize"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /account/favorites [GET]
// @Security ApiKeyAuth
func Favorites(ctx *gin.Context) {
	res := base.PageResponse{}
	pgnumber := ctx.Query("num")
	pgsize := ctx.Query("pagesize")
	Userid := "mingzheliu-ust-hk"
	fmt.Println(pgnumber, pgsize)
	code, data, count := service.GetItems(Userid, 1, 2, "favorites")
	res.SetCode(code)
	res.SetData(data)
	res.SetCount(count)
	ctx.JSON(http.StatusOK, res)
}

// DeleteItem @Description  DeleteItem: delete one item of one user, UserId is parsed from token
// @Tags         account
// @param 		 item-id   query   string   true   "item-id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /account/delete-item [POST]
// @Security ApiKeyAuth
func DeleteItem(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	ItemId := ctx.Query("item-id")
	code := service.DeleteItem(ItemId, UserId)
	res.SetCode(code)
	ctx.JSON(http.StatusOK, res)
}

// CreateCollectionByAccount @Description  create-collection: create one collection, only owner itself could create his/her collection, also parse UserID
// @Tags         account
// @param 		 name   query   string    true    "name"
// @param 		 logo-image   query   string    true    "logo-image"
// @param 		 feature-image  query   string    true    "feature-image"
// @param 		 banner-image   query   string    true    "banner-image"
// @param 		 description   query   string    true    "description"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /account/create-collection [POST]
// @Security ApiKeyAuth
func CreateCollectionByAccount(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"

	name := ctx.Query("name")
	logoImage := ctx.Query("logo-image")
	featureImage := ctx.Query("feature-image")
	bannerImage := ctx.Query("banner-image")

	description := ctx.Query("description")
	code, data := service.CreateCollectionByAccount(UserId, name, logoImage, featureImage, bannerImage, description)
	res.SetCode(code)
	res.SetData(data)

	ctx.JSON(http.StatusOK, res)
}

// EditCollection @Description  EditCollection: Edit one collection, only owner itself could edit his/her collection, also parse UserID, and do not require all params but lack param check
// @Tags         account
// @param 		 name   query   string    false    "name"
// @param 		 logo-image   query   string    false    "logo-image"
// @param 		 feature-image  query   string    false    "feature-image"
// @param 		 banner-image   query   string    false    "banner-image"
// @param 		 description   query   string    false    "description"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /account/edit-collection [POST]
// @Security ApiKeyAuth
func EditCollection(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"

	name := ctx.Query("name")
	logoImage := ctx.Query("logo-image")
	featureImage := ctx.Query("feature-image")
	bannerImage := ctx.Query("banner-image")
	description := ctx.Query("description")
	CollectionId := ctx.Query("collection-id")

	code, data := service.EditCollection(UserId, CollectionId, name, logoImage, featureImage, bannerImage, description)
	res.SetCode(code)
	res.SetData(data)
	ctx.JSON(http.StatusOK, res)
}

// Creation @Description  creation: return the collections created by one user, parse UserId from token
// @Tags         account
// @param 		 num   query   string   true   "num"
// @param 		 pagesize   query   string   true   "pagesize"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.PageResponse "error code and message and nil data"
// @Router       /account/creation [GET]
// @Security ApiKeyAuth
func Creation(ctx *gin.Context) {
	res := base.PageResponse{}
	pgnumber := ctx.Query("num")
	pgsize := ctx.Query("pagesize")
	Userid := "mingzheliu-ust-hk"
	fmt.Println(pgnumber, pgsize)
	code, data, count := service.GetCollections(Userid, 1, 2, "creation")
	res.SetCode(code)
	res.SetData(data)
	res.SetCount(count)
	ctx.JSON(http.StatusOK, res)
}

// DeleteCollection @Description  delete-collection: owner could delete his/her collection, parse UserId
// @Tags         account
// @param 		 collection-id   path   string    true    "collection-id"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /account/delete-collection [POST]
// @Security ApiKeyAuth
func DeleteCollection(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	CoId := ctx.Query("collection-id")
	code := service.DeleteCollection(CoId, UserId)
	res.SetCode(code)
	ctx.JSON(http.StatusOK, res)
}

// EditProfile @Description  edit-profile: owner edit his/her profile, do not require all params, parse UserId
// @Tags         account
// @param 		 banner-image   query   string    false    "banner-image"
// @param 		 avatar-image   query   string    false    "avatar-image"
// @param 		 poison  query   string    false    "poison"
// @param 		 campus   query   string    false    "campus"
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /account/edit-profile [POST]
// @Security ApiKeyAuth
func EditProfile(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	BannerImage := ctx.Query("banner-image")
	AvatarImage := ctx.Query("avatar-image")
	Poison := ctx.Query("poison")
	Campus := ctx.Query("campus")
	code, data := service.EditProfile(UserId, BannerImage, AvatarImage, Poison, Campus)
	res.SetCode(code)
	res.SetData(data)
	ctx.JSON(http.StatusOK, res)
}

// GetUser @Description  get-user: get account info of one user
// @Tags         account
// @Accept       json
// @Produce      json
// @Success 0 {object} base.ErrCode "Operation Succeed, code: 0"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /account/get-user [GET]
// @Security ApiKeyAuth
func GetUser(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	// 参数为id
	// 改成两种情况的返回：如id提供就是读其他人profile，如果不提供就是parse token
	code, data := service.GetUser(UserId)
	res.SetCode(code)
	res.SetData(data)
	ctx.JSON(http.StatusOK, res)
}
