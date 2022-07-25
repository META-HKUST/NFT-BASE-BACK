package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/sdk/service"
	"NFT-BASE-BACK/utils"
	"crypto/md5"
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
	"log"
	"strings"
)

var (
	verifier = emailverifier.NewVerifier()
)

func RegisterEmailToken(p model.Person, ReceiverName string) base.ErrCode {
	// check input
	if p.Email == "" {
		return base.EmptyInput
	}
	p.Token = utils.GenEmailToken()
	p.GenTime = utils.GetTimeNow()
	if err := p.StoreEmailToken(); err != nil {
		log.Println(err)
		return base.StoreEmailTokenError
	}
	if err := utils.Email(ReceiverName, p.Email, p.Token); err != nil {
		log.Println(err)
		return base.SendEmailError
	}
	return base.Success
}

func ActivateToken(token string) base.ErrCode {
	// check input
	if token == "" {
		return base.EmptyInput
	}
	_, genTime := model.GetGenTime(token)

	if genTime == "" {
		return base.TokenNotExist
	}
	if _, s := model.GetTokenStatus(token); s == "yes" {
		return base.TokenAlreadyActivated
	}
	if !utils.IsTimeValid(genTime) {
		return base.TokenInvalidError
	}
	if err := model.ActivateEmailToken(token); err != nil {
		log.Println(err)
		return base.ActivateEmailError
	} else {
		return base.Success
	}
}

func CheckEmailToken(p model.Person) base.ErrCode {
	// check input
	if p.Email == "" {
		return base.EmptyInput
	}
	p1, err := model.GetPerson(p.Email)
	if err != nil {
		log.Println(err)
		return base.GetPersonError
	}
	if p1.Activated == "no" {
		return base.TokenNotActivated
	}
	return base.Success
}

func Register(p model.Person) base.ErrCode {

	// check input
	if p.Email == "" {
		return base.EmptyInput
	}
	if p.Passwd == "" {
		return base.EmptyInput
	}
	//p.ActivateEmailToken()

	// check email format
	ret, err := verifier.Verify(p.Email)
	if err != nil {
		return base.EmailFormatError
	}
	if !ret.Syntax.Valid {
		return base.EmailFormatError
	}
	//// check if using ust email
	//b1 := strings.Contains(p.Email, "ust.hk")
	//if b1 == false{
	//	return base.EmailFormatError
	//}

	// check passwd length
	//l3 := strings.Count(p.Passwd, "") - 1
	//if l3 >= 24 || l3 < 8 {
	//	return base.PasswdLengthError
	//}

	// generate UserId
	t1 := strings.Replace(p.Email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)

	p1, _ := model.GetPerson(p.Email)
	log.Println(p1)
	if p1.Email == p.Email && p1.Activated == "no" {
		log.Println("delete account,profile,collection of : ", p.Email)
		_ = model.DeleteUser(p.Email)
		_ = model.DeleteProfile(p.Email)
		_ = model.DeleteCollection(UserId)
	} else if p1.Email == p.Email && p1.Activated == "yes" {
		return base.AccountExistError
	}

	//
	//if p1.Email == p.Email {
	//	if p1.Activated == "yes" {
	//		log.Println(base.AccountExistError.String())
	//		return base.AccountExistError
	//	} else {
	//		model.DeleteAccount(p1.Email)
	//	}
	//}

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(p.Passwd))
	Result := Md5Inst.Sum([]byte(""))

	p.Passwd = fmt.Sprintf("%x", Result)

	if err := p.Register(); err != base.Success {
		log.Println(err)
		return base.ServerError
	}
	log.Println("Register Succeed, email:", p.Email)
	if err := RegisterEmailToken(p, p.Email); err != base.Success {
		log.Println(err)
		return base.InputError
	}

	e1 := model.UpdateId(p.Email, UserId)
	if e1 != nil {
		log.Println(e1)
		return base.ServerError
	}
	e2 := model.InsertAccount(p.Email, UserId)
	if e2 != nil {
		log.Println(e2)
		return base.InsertProfileError
	}

	log.Println("Successfully insert into default profile: ", UserId)
	e3 := model.CreatCollection("Default Collection", "https://unift-1312994969.cos.ap-guangzhou.myqcloud.com/unifit/logo2.jpg?q-sign-algorithm=sha1&q-ak=AKIDBD4i9ML5aswlLgmfJisnTt30f6JJ6duu&q-sign-time=1658715913%3B1745115913&q-key-time=1658715913%3B1745115913&q-header-list=host&q-url-param-list=&q-signature=8021ba72ac3722d61da3ed551e60db7c13831cb3", "https://unift-1312994969.cos.ap-guangzhou.myqcloud.com/unifit/feature2.jpg?q-sign-algorithm=sha1&q-ak=AKIDBD4i9ML5aswlLgmfJisnTt30f6JJ6duu&q-sign-time=1658715943%3B1745115943&q-key-time=1658715943%3B1745115943&q-header-list=host&q-url-param-list=&q-signature=0b05784204c0b84b1c216f90565253f8fb909a43", "https://unift-1312994969.cos.ap-guangzhou.myqcloud.com/unifit/banner2.png?q-sign-algorithm=sha1&q-ak=AKIDBD4i9ML5aswlLgmfJisnTt30f6JJ6duu&q-sign-time=1658715986%3B1745115986&q-key-time=1658715986%3B1745115986&q-header-list=host&q-url-param-list=&q-signature=f4c8637dada6ffd079ceb66481aa7cb3022d71a0", 0, "Default Collection", UserId, p.Email, utils.GetTimeNow())
	if e3 != nil {
		log.Println(e3)
		return base.CreateCollectionError
	}
	err3 := service.Enroll(UserId)
	if err3 != nil {
		// TODO: check why this enroll error happens and if it influence transactions
		log.Println("Register to Fabric failed: ", err3, " ", UserId)
		return base.Success
	}

	return base.Success
}

func Login(p model.Person) (base.ErrCode, string, string) {
	// check input
	if p.Email == "" {
		return base.EmptyInput, "", ""
	}
	if p.Passwd == "" {
		return base.EmptyInput, "", ""
	}

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(p.Passwd))
	Result := Md5Inst.Sum([]byte(""))

	p.Passwd = fmt.Sprintf("%x", Result)

	if err := p.Login(); err != base.Success {
		return err, "", ""
	}
	token, err := utils.GenToken(p)
	if err != nil {
		return base.ServerError, "", ""
	}
	if err := CheckEmailToken(p); err != base.Success {
		return err, "", ""
	}
	t1 := strings.Replace(p.Email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)
	return base.Success, token, UserId
}

func ForgetPasswd(email string) base.ErrCode {
	// check input
	if email == "" {
		return base.EmptyInput
	}

	c := utils.GenVerifyCode()
	err := model.UpdateVerifyCode(email, c)
	if err != nil {
		return base.ServerError
	}
	reName, _ := model.GetNameByEmail(email)
	if reName == "" {
		reName = "Sir/Madam"
	}
	fmt.Println("reName: ", reName)
	utils.ResetEmail(reName, email, c)
	return base.Success
}

func ResetPasswd(email string, code string, pd string) base.ErrCode {
	// check input
	if email == "" {
		return base.EmptyInput
	}
	c1, err1 := model.GetVerifyCode(email)
	if err1 != nil {
		return base.InputError
	}
	fmt.Println("Get verify code: ", c1)

	// TODO: Set the expire time of verify code
	if c1 != code {
		return base.WrongVerifyCode
	}

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(pd))
	Result := Md5Inst.Sum([]byte(""))

	pd = fmt.Sprintf("%x", Result)

	err := model.ResetUpdate(email, pd)
	if err != nil {
		return base.InputError
	}
	return base.Success
}
