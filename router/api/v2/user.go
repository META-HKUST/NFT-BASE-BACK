package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Err1000 struct {
	Code int    `json:"code" default:"1000" example:"1000"`
	Msg  string `json:"msg" default:"Input Error" example:"Input Error"`
}

type Err2000 struct {
	Code int    `json:"code" default:"2000"  example:"2000"`
	Msg  string `json:"msg" default:"Server Error"  example:"Server Error"`
}

type ModelResponse struct {
	Code int    `json:"code" example:"0" default:"0"`                                // 错误码
	Msg  string `json:"msg" example:"Operation Succeed" default:"Operation Succeed"` // 错误描述
	Data string `json:"data" example:"" default:""`                                  // 返回数据
}

type RegisterRequest struct {
	Email  string `json:"email" example:"mingzheliu@ust.hk" default:"mingzheliu@ust.hk"`
	Passwd string `json:"passwd" example:"Abcd123456" default:"Abcd123456"`
}

// Register @Description  register: upload name email and passwd to register an account
// @Tags         user
// @param 		 RequestParam   body  RegisterRequest    true    "邮箱和密码"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /user/register [POST]
func Register(ctx *gin.Context) {
	// TODO (@mingzhe): associate the certificate with user info
	p := model.Person{
		Email:  ctx.PostForm("email"),
		Passwd: ctx.PostForm("passwd"),
	}
	res := base.Response{}
	fmt.Println(p)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type RerunEmailRequest struct {
	Email string `json:"email" example:"mingzheliu@ust.hk" default:"mingzheliu@ust.hk"`
}

// Rerun_Email @Description  rerun-email: 重新发送激活邮件
// @Tags         user
// @param 		 RequestParam   body  RerunEmailRequest    true    "邮箱"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /user/rerun-email [POST]
func Rerun_Email(ctx *gin.Context) {
	Email := ctx.PostForm("email")
	res := base.Response{}
	fmt.Println(Email)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

// Activate @Description  activate: 激活相应的邮件
// @Tags         user
// @param 		 token  query  string  true  "Email token sent to users"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /user/activate [GET]
func Activate(ctx *gin.Context) {
	res := base.Response{}
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type LoginRequest struct {
	Email  string `json:"email" example:"mingzheliu@ust.hk" default:"mingzheliu@ust.hk"`
	Passwd string `json:"passwd" example:"Abcd123456" default:"Abcd123456"`
}

// Login @Description  login: 登录
// @Tags         user
// @param 		 RequestParam  body  LoginRequest  true  "登录用的邮箱和密码"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /user/login [POST]
func Login(ctx *gin.Context) {
	Email := ctx.PostForm("email")
	res := base.Response{}
	res.SetData(gin.H{"user_id": "mazhengwang-ust-hk", "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im1pbmd6aGVsaXVAdXN0LmhrIiwiUGFzc3dkIjoiMTIzNDU2IiwiVG9rZW4iOiIiLCJBY3RpdmF0ZWQiOiIiLCJHZW5UaW1lIjoiIiwiZXhwIjoxNjU1MzcxNjA1LjQyNTU2MjksImlzcyI6IlllenppIn0.ULpqRvjdxyEd-B0FJHJMpg7-sIXb53JSqahRXDkYLPA"})
	fmt.Println(Email)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type Update_PasswdRequest struct {
	Old_Passwd string `json:"old_passwd" example:"Abcd123456" default:"Abcd123456"`
	New_Passwd string `json:"new_passwd" example:"Abcd1234567" default:"Abcd1234567"`
}

// Update_Passwd @Description  update-passwd: 登录状态下重置密码
// @Tags         user
// @param 		 RequestParam  body  Update_PasswdRequest  true  "登录用的邮箱和密码"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /user/update-passwd [POST]
func Update_Passwd(ctx *gin.Context) {
	res := base.Response{}
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type Forget_PasswdRequest struct {
	Email string `json:"email" example:"mingzheliu@ust.hk" default:"mingzheliu@ust.hk"`
}

// Forget_Passwd @Description  forget-passwd: 忘记密码，发送验证码
// @Tags         user
// @param 		 RequestParam  body  Forget_PasswdRequest  true  "登录用的邮箱和密码"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /user/forget-passwd [POST]
func Forget_Passwd(ctx *gin.Context) {
	res := base.Response{}
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type Reset_PasswdRequest struct {
	Email      string `json:"email" example:"mingzheliu@ust.hk" default:"mingzheliu@ust.hk"`
	Code       string `json:"code" example:"456WER" default:"456WER"`
	New_Passwd string `json:"Abcd12345" example:"Abcd12345" default:"Abcd12345"`
}

// Reset_Passwd @Description  reset_passwd: 输入邮箱、验证码和密码，重新设置已经忘记的密码
// @Tags         user
// @param 		 RequestParam  body  Reset_PasswdRequest  true  "邮箱、验证码密码"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /user/reset_passwd [POST]
func Reset_Passwd(ctx *gin.Context) {
	res := base.Response{}
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

type Edit_ProfileRequest struct {
	User_Name    string `json:"user_name" example:"Hunter" default:"Hunter"`
	Poison       string `json:"poison" example:"teacher" default:"teacher"`
	Organization string `json:"organization" example:"HKUST-GZ" default:"HKUST-GZ"`
}

type UserInfo struct {
	UserId           string `json:"user_id" `
	UserEmail        string `json:"user_email" `
	UserName         string `json:"user_name" `
	BannerImage      string `json:"banner_image" `
	LogoImage        string `json:"logo_image"`
	Poison           string `json:"poison" `
	Organization     string `json:"organization" `
	RegistrationTime string `json:"registration_time" `
}

func NewUserInfo() UserInfo {
	return UserInfo{
		UserId:           "mingzheliu-ust-hk",
		UserEmail:        "mingzheliu@ust.hk",
		UserName:         "LMZ",
		BannerImage:      "https://img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2Fsvc_rQkHVGf3aMI14v3pN-ZTI7uDRwN-QayvixX-nHSMZBgb1L1LReSg1-rXj4gNLJgAB0-yD8ERoT-Q2Gu4cy5AuSg-RdHF9bOxFDw%3Ds10000?fit=max&h=2500&w=2500&auto=format&s=61a1f05fd1f4a891c9b8fc197befc0a",
		LogoImage:        "img-ae.seadn.io/https%3A%2F%2Flh3.googleusercontent.com%2F7B0qai02OdHA8P_EOVK672qUliyjQdQDGNrACxs7WnTgZAkJa_wWURnIFKeOh5VTf8cfTqW3wQpozGedaC9mteKphEOtztls02RlWQ%3Ds10000?fit=max&h=120&w=120&auto=format&s=65b159799dcff448deaf9106b1ead13e",
		Poison:           "teacher",
		Organization:     "HKUST-GZ",
		RegistrationTime: "2022-06-16 20:45:40",
	}
}

// Edit_Profile @Description  edit-profile: 编辑用户的个人资料
// @Tags         user
// @param 		 logo_image   formData  file  false    "logo_image of a user"
// @param 		 banner_image formData  file  false    "banner_image of a user"
// @param 		 RequestParam  body  Edit_ProfileRequest  false  "用户名称、组织名称、老师还是学生"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Security ApiKeyAuth
// @Router       /user/edit-profile [POST]
func Edit_Profile(ctx *gin.Context) {
	res := base.Response{}
	res.SetData(NewUserInfo())
	res.SetCode(base.Success)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}

// GetUserInfo @Description  info: 获取用户的个人资料
// @Tags         user
// @param 		 user_id   query  string  true    "user_id"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /user/info [GET]
func GetUserInfo(ctx *gin.Context) {
	res := base.Response{}
	res.SetData(NewUserInfo())
	res.SetCode(base.Success)
	ctx.JSON(http.StatusOK, res.SetCode(base.Success))
}
