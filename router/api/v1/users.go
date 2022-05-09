package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Campus      int    `json:"campus" example:"1 mean CWB and 2 mean GZ"`
	Email       string `json:"email" example:"Sam@ust.hk"`
	Passwd      string `json:"passwd" example:"123"`
	BannerImage string `json:"bannerimage" example:"/home/yezzi/bannerimage"`
	AvatarImage string `json:"avatarimage" example:"/home/yezzi/bannerimage"`
	UserName    string `json:"username" example:"Sam"`
	Id          string `json:"id" example:"1001"`
	Certificate string `json:"certificate" example:"/home/yezzi/certificate_yezzi"`
}

// Register @Description  user register: upload all the parameters needed and get a success feedback
// @Tags         user
// @param 		 Campus        query   int       true    "campus"
// @param 		 Email         query   string    true    "email"
// @param 		 Passwd        query   string    true    "passwd"
// @param 		 BannerImage   query   file      true    "bannerimage"
// @param 		 AvatarImage   query   file      true    "avatarimage"
// @param 		 UserName      query   string    true    "username"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/users"
// @Router       /users/register [POST]
func Register(ctx *gin.Context) {
	// TODO (@mingzhe): store user info into db and complete related login function
	// TODO (@mingzhe): store pictures
	// TODO (@mingzhe): associate the certificate with user info
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
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
	// TODO (@mingzhe): verify the user info using db
	ctx.JSON(http.StatusOK, ctx.Request.Method+ctx.Request.URL.Path)
}
