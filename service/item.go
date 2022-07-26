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

func UpdateItem(token_id,ipfs_url string) (model.ItemInfo,base.ErrCode) {
	err := model.UpdateItem(token_id,ipfs_url)
	if err != nil{
		log.Println(err)
		return model.ItemInfo{},base.EditItemError
	}

	itemInfo,err := model.GetItemInfo(token_id)
	if err != nil{
		return model.ItemInfo{},base.GetItemError
	}
	return itemInfo,base.Success
}




