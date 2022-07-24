package service

import "NFT-BASE-BACK/model"

func GetMetaInfo(tokenId string) (model.MetaData,error){

	Info,err := model.GetMetaDataInfo(tokenId)
	if err != nil {
		return model.MetaData{},err
	}
	return Info,nil
}
