package v1

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register @Description  user register: upload all the parameters needed and get a success feedback
// @Tags         user
// @param 		 Email         query   string    true    "email"
// @param 		 Passwd        query   string    true    "passwd"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/register [POST]
func Register(ctx *gin.Context) {
	// TODO (@mingzhe): associate the certificate with user info
	p := model.Person{
		ctx.Query("email"),
		ctx.Query("passwd"),
	}
	if err := p.Register(); err != base.Success {
		ctx.JSON(http.StatusOK, gin.H{
			"code": base.AccountExistError,
			"msg":  base.AccountExistError.String(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": base.Success,
		"msg":  base.Success.String(),
	})
}

// Login @Description  user login
// @Tags         user
// @param 		 email      query   string    true    "email"
// @param 		 password   query   string    true    "password"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/login [POST]
func Login(ctx *gin.Context) {
	p := model.Person{
		ctx.Query("email"),
		ctx.Query("passwd"),
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
// @param 		 newpassword   query   string    true    "newpassword"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/login [POST]
func Update(ctx *gin.Context) {

	p := model.Person{
		ctx.Query("email"),
		ctx.Query("passwd"),
	}
	newpasswd := ctx.Query("newpasswd")
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
