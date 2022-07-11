package service

import "NFT-BASE-BACK/model"

func GetUserListByKeyWord(keyword string,pageNum,pageSize int64,) ([]model.UserProfile,error){
	userList,err :=model.GetUserListByKey(keyword,pageNum,pageSize)
	if err != nil {
		return []model.UserProfile{},err
	}
	return userList,nil
}
