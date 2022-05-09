package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description  user enroll
// @Tags         user
// @param 		 email   query   string    true    "email"
// @param 		 password   query   string    true    "password"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users [POST]
func Name21(ctx *gin.Context) {

	///  db

	///
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}

// @Description  user login
// @Tags         user
// @param 		 email   query   string    true    "email"
// @param 		 password   query   string    true    "password"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/login [POST]
func Name22(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}
