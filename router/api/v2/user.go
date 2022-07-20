package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/service"
	"crypto/md5"
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
	p := model.Person{}
	ctx.BindJSON(&p)
	res := base.Response{}
	code := service.Register(p)
	res.SetCode(code)
	ctx.JSON(http.StatusOK, res)
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
	p := model.Person{}
	ctx.BindJSON(&p)
	res := base.Response{}
	code := service.RegisterEmailToken(p, p.Email)
	ctx.JSON(http.StatusOK, res.SetCode(code))
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
	token := ctx.Query("token")
	code := service.ActivateToken(token)
	if code == base.Success {
		ctx.Redirect(http.StatusMovedPermanently, "https://unifit.ust.hk/register/success")
		return
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "  - https://unifit.ust.hk/register/fail")
	}
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
	p := model.Person{}
	ctx.BindJSON(&p)

	res := base.Response{}
	code, token, UserId := service.Login(p)
	if code != base.Success {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}
	res.SetCode(code)
	dto := gin.H{
		"token":   token,
		"user_id": UserId,
	}
	res.SetData(dto)
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
// @Security ApiKeyAuth
// @Router       /user/update-passwd [POST]
func Update_Passwd(ctx *gin.Context) {

	s, _ := ctx.Get("email")
	email := fmt.Sprintf("%v", s)
	ch := Update_PasswdRequest{}
	ctx.BindJSON(&ch)

	P := model.Person{
		Email:  email,
		Passwd: ch.Old_Passwd,
	}
	res := base.Response{}

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(ch.New_Passwd))
	Result := Md5Inst.Sum([]byte(""))

	ch.New_Passwd = fmt.Sprintf("%x", Result)

	code := P.Update(ch.New_Passwd)
	ctx.JSON(http.StatusOK, res.SetCode(code))
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

	ch := Forget_PasswdRequest{}
	ctx.BindJSON(&ch)
	code := service.ForgetPasswd(ch.Email)

	ctx.JSON(http.StatusOK, res.SetCode(code))
}

type Reset_PasswdRequest struct {
	Email      string `json:"email" example:"mingzheliu@ust.hk" default:"mingzheliu@ust.hk"`
	Code       string `json:"code" example:"466568" default:"456WER"`
	New_Passwd string `json:"new_Passwd" example:"Abcd12345" default:"Abcd12345"`
}

// Reset_Passwd @Description  reset_passwd: 输入邮箱、验证码和密码，重新设置已经忘记的密码
// @Tags         user
// @param 		 RequestParam  body  Reset_PasswdRequest  true  "邮箱、验证码密码"
// @Accept       json
// @Produce      json
// @Success 200  {object}   ModelResponse "Operation Succeed, code: 0 More details please refer to https://elliptic.larksuite.com/wiki/wikusjnG1KzGnrpQdmzjlqxDQVf"
// @Failure 400  {object}   Err1000       "Input error"
// @Failure 500  {object}   Err2000       "Server error"
// @Router       /user/reset-passwd [POST]
func Reset_Passwd(ctx *gin.Context) {
	res := base.Response{}

	ch := Reset_PasswdRequest{}
	ctx.BindJSON(&ch)
	code := service.ResetPasswd(ch.Email, ch.Code, ch.New_Passwd)

	ctx.JSON(http.StatusOK, res.SetCode(code))
}

type Edit_ProfileRequest struct {
	User_Name    string `json:"user_name" example:"Hunter" default:"Hunter"`
	Organization string `json:"organization" example:"HKUST-GZ" default:"HKUST-GZ"`
	Poison       string `json:"poison" example:"teacher" default:"teacher"`
	LogoImage    string `json:"logo_image" example:"https://unifit-1311571861.cos.ap-guangzhou.myqcloud.com/unifit/nft.jpg?q-sign-algorithm=sha1&q-ak=AKIDRikVzB8oDKBm68tOcYDcka9RSDhurYx5&q-sign-time=1656428492%3B1656432092&q-key-time=1656428492%3B1656432092&q-header-list=host&q-url-param-list=&q-signature=949835db0f086df54adc09d6e53dde318a74c2b6" default:"https://unifit-1311571861.cos.ap-guangzhou.myqcloud.com/unifit/nft.jpg?q-sign-algorithm=sha1&q-ak=AKIDRikVzB8oDKBm68tOcYDcka9RSDhurYx5&q-sign-time=1656428492%3B1656432092&q-key-time=1656428492%3B1656432092&q-header-list=host&q-url-param-list=&q-signature=949835db0f086df54adc09d6e53dde318a74c2b6"`
	LogoSig      string `json:"logo_image_signature" example:"abc" default:"teacher"`
	BannerImage  string `json:"banner_image" example:"https://unifit-1311571861.cos" default:"teacher"`
	BannerSig    string `json:"banner_image_signature" example:"abc" default:"teacher"`
	AvatarImage  string `json:"avatar_image" example:"https://unifit-1311571861.cos" default:"teacher" `
	AvatarSig    string `json:"avatar_image_signature" example:"abc" default:"teacher"`
}

type UserProfileInfo struct {
	UserId           string `json:"user_id"  example:"user_id"`
	UserEmail        string `json:"user_email"  example:"Hunter"`
	UserName         string `json:"user_name"  example:"Hunter"`
	BannerImage      string `json:"banner_image"  example:"Hunter"`
	LogoImage        string `json:"logo_image"  example:"Hunter"`
	Poison           string `json:"poison"  example:"Hunter"`
	Organization     string `json:"organization"  example:"Hunter"`
	RegistrationTime string `json:"registration_time"  example:"Hunter"`
}

// Edit_Profile @Description  edit-profile: 编辑用户的个人资料
// @Tags         user
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
	req := Edit_ProfileRequest{}

	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, "auth email error")
		return
	}

	ctx.BindJSON(&req)
	fmt.Println(req)
	userinfo, code := model.EditProfile(email.(string), req.User_Name, req.Organization, req.Poison, req.LogoImage, req.BannerImage)
	if code != base.Success {
		ctx.JSON(http.StatusInternalServerError, "Failed to edit information")
		return
	}
	res.Data = userinfo
	res.Code = 0
	ctx.JSON(http.StatusOK, res)
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
	userID := ctx.Query("user_id")
	//email, ok := ctx.Get("email")
	//if !ok {
	//	ctx.JSON(http.StatusInternalServerError, "auth email error")
	//	return
	//}
	code, userProfileInfo := model.GetUserInfoByID(userID)
	if userProfileInfo.LogoImage == "default image" {
		userProfileInfo.LogoImage = ""
	}
	if userProfileInfo.BannerImage == "default image" {
		userProfileInfo.BannerImage = ""
	}
	if userProfileInfo.Poison == "not set up" {
		userProfileInfo.Poison = ""
	}
	if userProfileInfo.Organization == "not set up" {
		userProfileInfo.Organization = ""
	}

	if code != base.Success {
		ctx.JSON(http.StatusOK, res.SetCode(base.ServerError))
		return
	}

	res.SetData(userProfileInfo)
	res.SetCode(base.Success)
	res.Msg = "Operation succeed"
	ctx.JSON(http.StatusOK, res)
}
