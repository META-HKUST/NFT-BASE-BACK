package v1

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/service"
	"NFT-BASE-BACK/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register @Description  user register: upload all the parameters needed and get a success feedback
// @Tags         user
// @param 		 Email         query   string    true    "email"
// @param 		 Passwd        query   string    true    "passwd"
// @param 		 Name        query   string    true    "name"
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
	//p.ActivateEmailToken()
	p1, _ := model.GetPerson(p.Email)

	if p1.Email == p.Email {
		if p1.Activated != "no" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": base.AccountExistError,
				"msg":  base.AccountExistError.String(),
			})
			return
		}
	}

	if err := p.Register(); err != base.Success {
		ctx.JSON(http.StatusOK, gin.H{
			"code": err,
			"msg":  err.String(),
		})
		return
	}
	name := ctx.Query("name")
	if err := service.RegisterEmailToken(p, name); err != base.Success {
		ctx.JSON(http.StatusOK, gin.H{
			"code": err,
			"msg":  err.String(),
		})
		return
	}
	ctx.JSON(http.StatusOK, new(base.Response).SetCode(base.Success))
}

// Login @Description  user login
// @Tags         user
// @param 		 email      query   string    true    "email"
// @param 		 passwd   query   string    true    "passwd"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/login [POST]
func Login(ctx *gin.Context) {
	p := model.Person{
		Email:  ctx.Query("email"),
		Passwd: ctx.Query("passwd"),
	}
	if err := p.Login(); err != base.Success {
		ctx.JSON(http.StatusOK, gin.H{
			"code": err,
			"msg":  err.String(),
		})
		return
	}
	token, err := utils.GenToken(p)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": base.GenTokenError,
			"msg":  err,
		})
		return
	}
	if err := service.CheckEmailToken(p); err != base.Success {
		ctx.JSON(http.StatusOK, gin.H{
			"code": err,
			"msg":  err.String(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": base.Success,
		"msg":  base.Success.String(),
		"data": gin.H{"token": token},
	})
}

// Update @Description  user password update
// @Tags         user
// @param 		 email      query   string    true    "email"
// @param 		 password   query   string    true    "password"
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
	fmt.Println(newpasswd)
	if err := p.Update(newpasswd); err != base.Success {
		ctx.JSON(http.StatusOK, gin.H{
			"code": err,
			"msg":  err.String(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": base.Success,
		"msg":  base.Success.String(),
	})
}

// Activate @Description  user email token activate
// @Tags         user
// @param 		 token    query   string    true    "token"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/activate [POST]
func Activate(ctx *gin.Context) {

	token := ctx.Query("token")
	err := service.ActivateToken(token)
	if err != base.Success {
		ctx.JSON(http.StatusOK, gin.H{
			"code": err,
			"msg":  err.String(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": base.Success,
		"msg":  base.Success.String(),
	})
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
	//if err := service.RegisterEmailToken(p, name); err != base.Success {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code": err,
	//		"msg":  err.String(),
	//	})
	//	return
	//}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": base.Success,
	//	"msg":  base.Success.String(),
	//})
}
