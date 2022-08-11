package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListReq struct {
	Email string `json:"email" example:"mingzheliu@ust.hk"`
}

// AddWhiteList
// @Tags         userList
// @param 		 RequestParam  body  UserListReq  false  "邮箱"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /userList/add_white [POST]
func AddWhiteList(ctx *gin.Context) {

	res := base.Response{}
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, res.SetCode(base.InputError))
	}
	if email != "admin@unifit.art" && email != "mingzheliu@ust.hk" && email != "unifit@hkust-gz.edu.cn" {
		ctx.JSON(http.StatusOK, res.SetCode(base.PermissionDenied))
		return
	}

	var g UserListReq
	ctx.BindJSON(&g)

	err := model.AddWhiteList(g.Email)

	if err != nil {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetCode(base.Success)
	ctx.JSON(http.StatusOK, res.SetData("WhiteList added: "+g.Email))

}

// AddBlackList
// @Tags         userList
// @param 		 RequestParam  body  UserListReq  false  "邮箱"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /userList/add_black [POST]
func AddBlackList(ctx *gin.Context) {
	res := base.Response{}
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, res.SetCode(base.InputError))
	}
	if email != "admin@unifit.art" && email != "mingzheliu@ust.hk" && email != "unifit@hkust-gz.edu.cn" {
		ctx.JSON(http.StatusOK, res.SetCode(base.PermissionDenied))
		return
	}

	var g UserListReq
	ctx.BindJSON(&g)

	err := model.AddBlackList(g.Email)

	if err != nil {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetCode(base.Success)
	ctx.JSON(http.StatusOK, res.SetData("BlackList added: "+g.Email))
}

// DeleteWhiteList
// @Tags         userList
// @param 		 RequestParam  body  UserListReq  false  "邮箱"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /userList/delete_white [POST]
func DeleteWhiteList(ctx *gin.Context) {
	res := base.Response{}
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, res.SetCode(base.InputError))
	}
	if email != "admin@unifit.art" && email != "mingzheliu@ust.hk" && email != "unifit@hkust-gz.edu.cn" {
		ctx.JSON(http.StatusOK, res.SetCode(base.PermissionDenied))
		return
	}

	var g UserListReq
	ctx.BindJSON(&g)

	err := model.DeleteWhite(g.Email)

	if err != nil {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetCode(base.Success)
	ctx.JSON(http.StatusOK, res.SetData("WhiteList delete: "+g.Email))
}

// DeleteBlackList
// @Tags         userList
// @param 		 RequestParam  body  UserListReq  false  "邮箱"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /userList/delete_black [POST]
func DeleteBlackList(ctx *gin.Context) {
	res := base.Response{}
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, res.SetCode(base.InputError))
		return
	}
	if email != "admin@unifit.art" && email != "mingzheliu@ust.hk" && email != "unifit@hkust-gz.edu.cn" {
		ctx.JSON(http.StatusOK, res.SetCode(base.PermissionDenied))
		return
	}

	var g UserListReq
	ctx.BindJSON(&g)

	err := model.DeleteBlack(g.Email)

	if err != nil {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetCode(base.Success)
	ctx.JSON(http.StatusOK, res.SetData("BlackList delete: "+g.Email))
}

// SearchWhiteList
// @Tags         userList
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /userList/search_white [GET]
func SearchWhiteList(ctx *gin.Context) {
	res := base.Response{}
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, res.SetCode(base.InputError))
		return
	}
	if email != "admin@unifit.art" && email != "mingzheliu@ust.hk" && email != "unifit@hkust-gz.edu.cn" {
		ctx.JSON(http.StatusOK, res.SetCode(base.PermissionDenied))
		return
	}

	l, err := model.GetWhiteList()

	if err != nil {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetCode(base.Success)
	ctx.JSON(http.StatusOK, res.SetData(l))
}

// SearchBlackList
// @Tags         userList
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /userList/search_black [GET]
func SearchBlackList(ctx *gin.Context) {
	res := base.Response{}
	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, res.SetCode(base.InputError))
		return
	}
	if email != "admin@unifit.art" && email != "mingzheliu@ust.hk" && email != "unifit@hkust-gz.edu.cn" {
		ctx.JSON(http.StatusOK, res.SetCode(base.PermissionDenied))
		return
	}

	l, err := model.GetBlackList()

	if err != nil {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetCode(base.Success)
	ctx.JSON(http.StatusOK, res.SetData(l))
}
