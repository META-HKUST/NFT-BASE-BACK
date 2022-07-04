package model

import (
	"NFT-BASE-BACK/base"
	"log"
)

var(
	queryBalanceByID = string("select token from accounts where user_id=?;")
	updateBalance = string("update accounts set token=? where user_id=?;")
)

func Transfer(tokenNum float32,fromUserId, toUserId string) base.ErrCode{
	var FromBalance float32
	var ToBalance	float32
	err := db.Get(&FromBalance,queryBalanceByID,fromUserId)
	if err != nil {
		log.Println(base.QueryError, base.QueryError.String(), err)
		return base.QueryError
	}
	if FromBalance < tokenNum {
		log.Println(base.BalanceNotEnough, base.BalanceNotEnough.String(), err)
		return base.BalanceNotEnough
	}
	err = db.Get(&ToBalance,queryBalanceByID,toUserId)
	if err != nil {
		log.Println(base.QueryError, base.QueryError.String(), err)
		return base.QueryError
	}
	FromBalance -= tokenNum
	ToBalance += tokenNum
	tx, err := db.Beginx()
	if err != nil {
		log.Println("Transaction start failed",err)
		return base.UpdateBalanceError
	}
	tx.MustExec(updateBalance,FromBalance,fromUserId)
	tx.MustExec(updateBalance,ToBalance,toUserId)
	err = tx.Commit()
	if err !=nil {
		log.Println(err)
		return base.UpdateBalanceError
	}
	return base.Success
}

func GetTokenInfo(userid string) (float32,base.ErrCode){
	var Balance	float32
	err := db.Get(&Balance,queryBalanceByID,userid)
	if err != nil {
		log.Println(base.QueryError, base.QueryError.String(), err)
		return Balance,base.QueryError
	}
	return Balance,base.Success
}
