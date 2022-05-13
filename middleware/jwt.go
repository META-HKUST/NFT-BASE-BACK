package middleware

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// JWTAuth middleware
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": base.LackTokenError,
				"msg":  base.LackTokenError.String(),
			})
			return
		}
		log.Print("token:", authHeader)

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": base.AuthFormatError,
				"msg":  base.AuthFormatError.String(),
			})
			ctx.Abort()
			return
		}

		//解析token包含的信息
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": base.InvalidToken,
				"msg":  base.InvalidToken.String(),
			})
			ctx.Abort()
			return
		}

		if err := CheckUserInfo(claims); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": base.UserTokenError,
				"msg":  base.UserTokenError.String(),
			})
			ctx.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		ctx.Set("claims", claims)
		ctx.Next() // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息

	}
}

//检查用户名信息
func CheckUserInfo(claims *utils.CustomClaims) error {
	p := model.Person{
		claims.Email,
		claims.Passwd,
	}
	if p.Login() == base.Success {
		return nil
	}
	return errors.New("Wrong passwd or email")
}
