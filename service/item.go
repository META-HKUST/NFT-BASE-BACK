package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"log"
)

func EditItem(itemId,itemName,description,collectionId string,label []string) (model.ItemInfo,base.ErrCode) {
	err := model.EditItem(itemId,itemName,description,collectionId,label)
	if err != nil{
		log.Println(err)
		return model.ItemInfo{},base.EditItemError
	}

	itemInfo,err := model.GetItemInfo(itemId)
	if err != nil{
		return model.ItemInfo{},base.GetItemError
	}
	return itemInfo,base.Success
}


