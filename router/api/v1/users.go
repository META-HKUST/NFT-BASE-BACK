package v1

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/service"
	"NFT-BASE-BACK/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register @Description  user register: upload all the parameters needed and get a success feedback
// @Tags         user
// @param 		 email         query   string    true    "email"
// @param 		 passwd        query   string    true    "passwd"
// @param 		 name        query   string    true    "name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/register [POST]
func Register(ctx *gin.Context) {
	// TODO (@mingzhe): associate the certificate with user info
	p := model.Person{
		Email:  ctx.Query("email"),
		Passwd: ctx.Query("passwd"),
	}
	res := base.Response{}

	//p.ActivateEmailToken()
	p1, _ := model.GetPerson(p.Email)

	if p1.Email == p.Email {
		if p1.Activated != "no" {
			ctx.JSON(http.StatusOK, res.SetCode(base.AccountExistError))
			return
		}
	}

	if err := p.Register(); err != base.Success {
		ctx.JSON(http.StatusOK, res.SetCode(err))
		return
	}
	name := ctx.Query("name")
	if err := service.RegisterEmailToken(p, name); err != base.Success {
		ctx.JSON(http.StatusOK, res.SetCode(err))
		return
	}
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

// Login @Description  user login
// @Tags         user
// @param 		 email      query   string    true    "email"
// @param 		 passwd   query   string    true    "password"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/login [POST]
func Login(ctx *gin.Context) {
	p := model.Person{
		Email:  ctx.Query("email"),
		Passwd: ctx.Query("passwd"),
	}
	res := base.Response{}
	if err := p.Login(); err != base.Success {
		ctx.JSON(http.StatusOK, res.SetCode(err))
		return
	}
	token, err := utils.GenToken(p)
	if err != nil {
		ctx.JSON(http.StatusOK, res.SetCode(base.GenTokenError))
		return
	}
	if err := service.CheckEmailToken(p); err != base.Success {
		ctx.JSON(http.StatusOK, res.SetCode(err))
		return
	}
	res.SetCode(base.Success)
	res.SetData(gin.H{"token": token, "id": "mazhengwang-ust-hk"})
	ctx.JSON(http.StatusOK, res)
}

// Update @Description  user password update
// @Tags         user
// @param 		 email      query   string    true    "email"
// @param 		 passwd   query   string    true    "password"
// @param 		 newpasswd   query   string    true    "newpasswd"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/update [POST]
func Update(ctx *gin.Context) {
	p := model.Person{
		Email:  ctx.Query("email"),
		Passwd: ctx.Query("passwd"),
	}
	newpasswd := ctx.Query("newpasswd")
	res := base.Response{}
	code := p.Update(newpasswd)
	ctx.JSON(http.StatusOK, res.SetCode(code))
}

// Activate @Description  user email token activate
// @Tags         user
// @param 		 token    query   string    true    "token"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/activate [POST]
func Activate(ctx *gin.Context) {
	res := base.Response{}
	token := ctx.Query("token")
	code := service.ActivateToken(token)
	ctx.JSON(http.StatusOK, res.SetCode(code))
}

// RerunEmail @Description  send activation email again
// @Tags         user
// @param 		 email    query   string    true    "email"
// @param 		 name    query   string    true    "name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/rerunEmail [POST]
func RerunEmail(ctx *gin.Context) {
	p := model.Person{
		Email: ctx.Query("email"),
	}
	name := ctx.Query("name")

	res := base.Response{}
	code := service.RegisterEmailToken(p, name)

	ctx.JSON(http.StatusOK, res.SetCode(code))
}

// ResetPasswd @Description  send activation email again
// @Tags         user
// @param 		 email    query   string    true    "email"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/rerunEmail [POST]
func ForgetPasswd(ctx *gin.Context) {
	p := model.Person{
		Email: ctx.Query("email"),
	}
	res := base.Response{}
	code := service.ForgetPasswd(p.Email)
	ctx.JSON(http.StatusOK, res.SetCode(code))
}

// ResetPasswd @Description  send activation email again
// @Tags         user
// @param 		 email    query   string    true    "email"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/rerunEmail [POST]
func ResetPasswd(ctx *gin.Context) {
	p := model.Person{
		Email:  ctx.Query("email"),
		Passwd: ctx.Query("passwd"),
	}
	verify := ctx.Query("verifycode")
	res := base.Response{}
	code := service.ResetPasswd(p.Email, p.Passwd, verify)

	ctx.JSON(http.StatusOK, res.SetCode(code))
}

// @Description  create-collection
// @Tags         account
// @param 		 id   path   string    true    "user id"
// @param 		 collection-Name   query   string   false   "collection Name"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/Anna/create-collection"
// @Router       /{id}/create-collection [POST]

func DeleteUUser(ctx *gin.Context) {
	res := base.Response{}
	ctx.JSON(http.StatusOK, res)
}
