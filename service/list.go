package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"log"
)

func GetItem(itemId string, userId string) (model.ItemInfoAndLike, base.ErrCode) {
	itemInfo, err := model.GetItemAndLikeInfo(itemId, userId)
	if err != nil {
		log.Println(err)
		return model.ItemInfoAndLike{}, base.QueryError
	}
	return itemInfo, base.Success

}

//func GetItemList(page_num, page_size int64, userId string, userLike, userCollect, userCreate bool, category string, keyword string, rank_favorite, rank_time bool, collection_id int) ([]model.ItemAndLogo, error) {
//	items, err := model.GetItemList(page_num, page_size, userId, userLike, userCollect, userCreate, category, keyword, rank_favorite, rank_time, collection_id)
//	if err != nil {
//		log.Println(err)
//		return []model.ItemAndLogo{}, err
//	}
//	log.Println("items: ", items)
//	return items, nil
//}
//
//func GetCollectionList(page_num, page_size int64, userId string, keyword string, rank_favorite, rank_time bool, label string) ([]model.Collection, error) {
//	collections, err := model.GetCollectionList(page_num, page_size, userId, keyword, rank_favorite, rank_time, label)
//	if err != nil {
//		log.Println(err)
//		return []model.Collection{}, err
//	}
//	return collections, nil
//}
