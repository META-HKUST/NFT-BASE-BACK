package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
)


type Token struct {
	UserId    string `json:"user_id"`
	Token	  float32 `json:"token"`
}


func Transfer(tokenNum float32,fromUserId, toUserId string) (Token,base.ErrCode){
	var balance float32

	errCode := model.Transfer(tokenNum,fromUserId,toUserId)
	if errCode != base.Success {
		return Token{},errCode
	}

	balance,errCode = model.GetTokenInfo(fromUserId)
	if errCode != base.Success {
		return Token{},errCode
	}


	return Token{UserId: fromUserId,Token: balance},errCode

}


func GetTokenInfo(email string) (Token,base.ErrCode){
	userID, err := model.GetUserIDByEmail(email)
	if err != nil {
		return Token{}, base.UserIDNotExist
	}

	balance,code := model.GetTokenInfo(userID)
	if code != base.Success {
		return Token{},code
	}
	return Token{userID,balance},code
}
