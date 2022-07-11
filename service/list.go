package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"log"
)

func GetItem(itemId string) (model.ItemInfo,base.ErrCode){
	itemInfo, err := model.GetItemInfo(itemId)
	if err != nil {
		log.Println(err)
		return model.ItemInfo{},base.QueryError
	}

	return itemInfo,base.Success

}

func GetItemList(page_num,page_size int64,userId string,userLike,userCollect,userCreate bool,category string,keyword string,rank_favorite,rank_time bool,collection_id int) ([]model.Item,error){
	items,err := model.GetItemList(page_num,page_size,userId,userLike,userCollect,userCreate,category,keyword,rank_favorite,rank_time,collection_id)
	if err != nil {
		log.Println(err)
		return []model.Item{},err
	}
	return items,nil
}
