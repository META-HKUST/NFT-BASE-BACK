package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/sdk/service"
	"NFT-BASE-BACK/utils"
	"crypto/md5"
	"fmt"
	"log"
	"strings"
)

func RegisterEmailToken(p model.Person, ReceiverName string) base.ErrCode {
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
	// TODO (@mingzhe): associate the certificate with user info

	//p.ActivateEmailToken()
	// check input
	p1, _ := model.GetPerson(p.Email)
	if p.Email == "" {
		return base.EmptyInput
	}
	if p.Passwd == "" {
		return base.EmptyInput
	}

	if p1.Email == p.Email {
		if p1.Activated != "no" {
			log.Println(base.AccountExistError.String())
			return base.AccountExistError
		}
	}

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(p.Passwd))
	Result := Md5Inst.Sum([]byte(""))

	p.Passwd = fmt.Sprintf("%x", Result)

	if err := p.Register(); err != base.Success {
		log.Println(err)
		return base.ErrCode(err)
	}
	name := "Sir/Madam"
	if err := RegisterEmailToken(p, name); err != base.Success {
		log.Println(err)
		return base.ErrCode(err)
	}

	t1 := strings.Replace(p.Email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)

	e1 := model.UpdateId(p.Email, UserId)
	if e1 != nil {
		log.Println(e1)
		return base.ServerError
	}
	e2 := model.InsertAccount(p.Email, UserId)
	if e2 != nil {
		log.Println(e2)
		return base.ServerError
	}
	err := service.Enroll(UserId)
	if err != nil {
		return base.EnrollFail
	}

	return base.Success
}

func Login(p model.Person) (base.ErrCode, string, string) {

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(p.Passwd))
	Result := Md5Inst.Sum([]byte(""))

	p.Passwd = fmt.Sprintf("%x", Result)

	if err := p.Login(); err != base.Success {
		return base.InputError, "", ""
	}
	token, err := utils.GenToken(p)
	if err != nil {
		return base.ServerError, "", ""
	}
	if err := CheckEmailToken(p); err != base.Success {
		return base.ServerError, "", ""
	}
	t1 := strings.Replace(p.Email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)
	return base.Success, token, UserId
}

func ForgetPasswd(emial string) base.ErrCode {
	c := utils.GenVerifyCode()
	err := model.UpdateVerifyCode(emial, c)
	if err != nil {
		return base.ServerError
	}
	utils.ResetEmail("Sir/Madam", emial, c)
	return base.Success
}

func ResetPasswd(email string, code string, pd string) base.ErrCode {
	c1, err1 := model.GetVerifyCode(email)
	if err1 != nil {
		return base.InputError
	}
	fmt.Println("Get verify code: ", c1)

	// TODO: Set the expire time of verify code
	if c1 != code {
		return base.WrongVerifyCode
	}
	err := model.ResetUpdate(email, pd)
	if err != nil {
		return base.InputError
	}
	return base.Success
}
