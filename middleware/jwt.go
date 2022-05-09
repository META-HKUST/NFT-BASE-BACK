package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware(ctx *gin.Context) {
	ctx.Next()
}
