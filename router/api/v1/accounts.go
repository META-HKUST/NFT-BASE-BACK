package v1

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description  create single item: parse UserId from token and create NFT(Creater and Owner are defined as UserId)
// @Tags         account
// @param 		 name   query   string   true   "name"
// @param 		 image   query   string   true    "image"
// @param 		 description   query   string   true    "description"
// @param 		 collection   query   string   true    "collection"
// @param 		 category   query   string   true    "category"
// @param 		 label   query   string   true    "label"
// @Accept       json
// @Produce      json
// @Success      200 {object} _Item
// @Router       /{id}/create-item [POST]
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

// @Description  edit single item: parse UserId from token and edit NFT
// @Tags         account
// @param 		 name   query   string   false   "name"
// @param 		 image   query   string   false   "image"
// @param 		 description   query   string   false   "description"
// @param 		 collection   query   string   false   "collection"
// @param 		 category   query   string   false   "category"
// @param 		 label   query   string   false   "label"
// @Accept       json
// @Produce      json
// @Success      200 {object} _Item
// @Router       /{id}/create-item [POST]
func EditItem(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	name := ctx.Query("name")
	image := ctx.Query("image")
	description := ctx.Query("description")
	itemCollection := ctx.Query("collection")
	category := ctx.Query("category")
	label := ctx.Query("label")
	code, data := service.EditItem(UserId, name, image, description, itemCollection, category, []string{label})
	res.SetCode(code)
	res.SetData(data)
	ctx.JSON(http.StatusOK, res)
}

// @Description  collected: will return the collected NFT items of one user, enter pagesize and pagenumber return number and info
// @Tags         account
// @param 		 pagenumber   query   string   true   "pagenumber"
// @param 		 pagesize   query   string   true   "pagesize"
// @Accept       json
// @Produce      json
// @Success      200 []{object} _[]Item
// @Router       /{id}/collected [GET]
func Collected(ctx *gin.Context) {
	res := base.PageResponse{}
	UserId := "mingzheliu-ust-hk"
	pgnumber := ctx.Query("pagenumber")
	pgsize := ctx.Query("pagesize")
	fmt.Println(pgnumber, pgsize)
	code, data, count := service.GetItems(UserId, 1, 2, "collected")
	res.SetCode(code)
	res.SetData(data)
	res.SetCount(count)
	ctx.JSON(http.StatusOK, res)
}

// @Description  favorites: will return the favorite NFT items of one user, in which UserId is parsed from token
// @Tags         account
// @param 		 pagenumber   query   string   true   "pagenumber"
// @param 		 pagesize   query   string   true   "pagesize"
// @Accept       json
// @Produce      json
// @Success      200 []{object} _[]Item
// @Router       /{id}/favorites [GET]
func Favorites(ctx *gin.Context) {
	res := base.PageResponse{}
	pgnumber := ctx.Query("pagenumber")
	pgsize := ctx.Query("pagesize")
	Userid := "mingzheliu-ust-hk"
	fmt.Println(pgnumber, pgsize)
	code, data, count := service.GetItems(Userid, 1, 2, "favorites")
	res.SetCode(code)
	res.SetData(data)
	res.SetCount(count)
	ctx.JSON(http.StatusOK, res)
}

// @Description  DeleteItem: delete one item of one user, UserId is parsed from token
// @Tags         account
// @param 		 item-id   query   string   true   "item-id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "..."
// @Router       /{id}/delete-item [POST]
func DeleteItem(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	ItemId := ctx.Query("item-id")
	code := service.DeleteItem(ItemId, UserId)
	res.SetCode(code)
	ctx.JSON(http.StatusOK, res)
}

// @Description  create-collection: create one collection, only owner itself could create his/her collection, also parse UserID
// @Tags         account
// @param 		 name   query   string    true    "name"
// @param 		 logo-image   query   string    true    "logo-image"
// @param 		 feature-image  query   string    true    "feature-image"
// @param 		 banner-image   query   string    true    "banner-image"
// @param 		 description   query   string    true    "description"
// @Accept       json
// @Produce      json
// @Success      200 {object} collection
// @Router       /{id}/create-collection [POST]
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

// @Description  EditCollection: Edit one collection, only owner itself could edit his/her collection, also parse UserID, and do not require all params but lack param check
// @Tags         account
// @param 		 name   query   string    false    "name"
// @param 		 logo-image   query   string    false    "logo-image"
// @param 		 feature-image  query   string    false    "feature-image"
// @param 		 banner-image   query   string    false    "banner-image"
// @param 		 description   query   string    false    "description"
// @Accept       json
// @Produce      json
// @Success      200 {object} collection
// @Router       /{id}/edit-collection [POST]
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

// @Description  creation: return the collections created by one user, parse UserId from token
// @Tags         account
// @param 		 pagenumber   query   string   true   "pagenumber"
// @param 		 pagesize   query   string   true   "pagesize"
// @Accept       json
// @Produce      json
// @Success      200 []{object} []collection
// @Router       /{id}/creation [GET]
func Creation(ctx *gin.Context) {
	res := base.PageResponse{}
	pgnumber := ctx.Query("pagenumber")
	pgsize := ctx.Query("pagesize")
	Userid := "mingzheliu-ust-hk"
	fmt.Println(pgnumber, pgsize)
	code, data, count := service.GetCollections(Userid, 1, 2, "creation")
	res.SetCode(code)
	res.SetData(data)
	res.SetCount(count)
	ctx.JSON(http.StatusOK, res)
}

// @Description  delete-collection: owner could delete his/her collection, parse UserId
// @Tags         account
// @param 		 collection-id   path   string    true    "collection-id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "..."
// @Router       /{id}/delete-collection [POST]
func DeleteCollection(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	CoId := ctx.Query("collection-id")
	code := service.DeleteCollection(CoId, UserId)
	res.SetCode(code)
	ctx.JSON(http.StatusOK, res)
}

// @Description  edit-profile: owner edit his/her profile, do not require all params, parse UserId
// @Tags         account
// @param 		 banner-image   query   string    false    "banner-image"
// @param 		 avatar-image   query   string    false    "avatar-image"
// @param 		 poison  query   string    false    "poison"
// @param 		 campus   query   string    false    "campus"
// @Accept       json
// @Produce      json
// @Success      200 {object} []Account
// @Router       /{id}/edit-profile [POST]
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

// @Description  get-user: get account info of one user
// @Tags         account
// @Accept       json
// @Produce      json
// @Success      200 {object} []Account
// @Router       /{id}/get-user [GET]
func GetUser(ctx *gin.Context) {
	res := base.Response{}
	UserId := "mingzheliu-ust-hk"
	code, data := service.GetUser(UserId)
	res.SetCode(code)
	res.SetData(data)
	ctx.JSON(http.StatusOK, res)
}
