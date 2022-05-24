package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/utils"
)

func RegisterEmailToken(p model.Person, ReceiverName string) base.ErrCode {
	p.Token = utils.GenEmailToken()
	p.GenTime = utils.GetTimeNow()
	if err := p.StoreEmailToken(); err != nil {
		return base.StoreEmailTokenError
	}
	if err := utils.Email(ReceiverName, p.Email, p.Token); err != nil {
		return base.SendEmailError
	}
	return base.Success
}

func ActivateToken(token string) base.ErrCode {
	genTime := model.GetGenTime(token)
	if genTime == "" {
		return base.TokenNotExist
	}
	if s := model.GetTokenStatus(token); s == "yes" {
		return base.TokenAlreadyActivated
	}
	if !utils.IsTimeValid(genTime) {
		return base.TokenInvalidError
	}
	if err := model.ActivateEmailToken(token); err != nil {
		return base.ActivateEmailError
	} else {
		return base.Success
	}
}

func CheckEmailToken(p model.Person) base.ErrCode {
	p1, err := model.GetPerson(p.Email)
	if err != nil {
		return base.GetPersonError
	}
	if p1.Activated == "no" {
		return base.TokenNotActivated
	}
	return base.Success
}
