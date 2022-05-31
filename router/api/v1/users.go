package v1

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/service"
	"NFT-BASE-BACK/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register @Description  register: upload name email and passwd to register an account
// @Tags         user
// @param 		 email       query   string    true    "email"
// @param 		 passwd      query   string    true    "passwd"
// @param 		 name        query   string    true    "the user's name"
// @Accept       json
// @Produce      json
// @Success 200 {object} base.Response "Operation Succeed"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
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

// Login @Description  login: enter the email and passwd and return token and UserId
// @Tags         user
// @param 		 email      query   string    true    "email"
// @param 		 passwd   query   string    true    "passwd"
// @Accept       json
// @Produce      json
// @Success 200 {object} base.Response "Operation Succeed"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
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

// Update @Description  update: password update
// @Tags         user
// @param 		 email      query   string    true    "email"
// @param 		 passwd   query   string    true    "passwd"
// @param 		 newpasswd   query   string    true    "newpasswd"
// @Accept       json
// @Produce      json
// @Success 200 {object} base.Response "Operation Succeed"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
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

// Activate @Description  activate: activate email using the Email token generated within 15 minutes
// @Tags         user
// @param 		 token    query   string    true    "token of email(different from auth jwt token)"
// @Accept       json
// @Produce      json
// @Success 200 {object} base.Response "Operation Succeed"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /users/activate [GET]
func Activate(ctx *gin.Context) {
	res := base.Response{}
	token := ctx.Query("token")
	code := service.ActivateToken(token)
	ctx.JSON(http.StatusOK, res.SetCode(code))
}

// RerunEmail @Description  rerun-email: send activation email to the email address again
// @Tags         user
// @param 		 email    query   string    true    "email"
// @param 		 name    query   string    true    "name"
// @Accept       json
// @Produce      json
// @Success 200 {object} base.Response "Operation Succeed"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /users/rerun-Email [POST]
func RerunEmail(ctx *gin.Context) {
	p := model.Person{
		Email: ctx.Query("email"),
	}
	name := ctx.Query("name")

	res := base.Response{}
	code := service.RegisterEmailToken(p, name)

	ctx.JSON(http.StatusOK, res.SetCode(code))
}

// ForgetPasswd @Description  forget-passwd: enter the email and send VerifyCode to change passwd
// @Tags         user
// @param 		 email    query   string    true    "email"
// @Accept       json
// @Success 200 {object} base.Response "Operation Succeed"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /users/forget-passwd [POST]
func ForgetPasswd(ctx *gin.Context) {
	p := model.Person{
		Email: ctx.Query("email"),
	}
	res := base.Response{}
	code := service.ForgetPasswd(p.Email)
	ctx.JSON(http.StatusOK, res.SetCode(code))
}

// ResetPasswd @Description  reset-passwd: send activation email again
// @Tags         user
// @param 		 email    query   string    true    "email"
// @Accept       json
// @Produce      json
// @Success 200 {object} base.Response "Operation Succeed"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /users/reset-passwd [POST]
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

// DeleteUUser @Description  delete-user: check User's authority and close one's account
// @Tags         user
// @param 		 UserId  query   string    true    "UserId"
// @Accept       json
// @Produce      json
// @Success 200 {object} base.Response "Operation Succeed"
// @Failure 400 {object} base.ErrCode "request error"
// @Failure 500 {object} base.Response "error code and message and nil data"
// @Router       /{id}/delete-user [POST]

func DeleteUser(ctx *gin.Context) {
	res := base.Response{}
	ctx.JSON(http.StatusOK, res)
}
