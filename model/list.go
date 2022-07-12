package model

import (
	"fmt"
	"log"
	"strconv"
)

var (

	queryByCondition = string("select * from items limit ? offset ?")
	queryconditions = string("select * from items ")
	queryByRankFavourite = string("select * from items order by rank_favourite desc limit ? offset ?")
	queryByRankTime = string("select * from items order by rank_time desc limit ? offset ?")
	queryCollections = string("select * from collection limit ? offset ?")
	collectionByCondition = string("select * from collection ")
)

type Collection struct {
	CollectionId string   `json:"collection_id" db:"collection_id"`
	CollectionName string   `json:"collection_name" db:"collection_name"`
	BannerImage    string   `json:"banner_image" db:"banner_image"`
	LogoImage      string   `json:"logo_image" db:"logo_image"`
	FeatureImage   string   `json:"feature_image" db:"feature_image"`
	Description    string   `json:"description" db:"description"`
	ItemNum        int      `json:"item_num" db:"items_count"`
	OwnerId        string   `json:"owner_id" db:"owner"`
	OwnerName      string   `json:"owner_name" db:"creater"`
	CreationTime   string   `json:"creation_time" db:"created_at"`
}

func GetItemList(page_num,page_size int64,userId string,userLike,userCollect,userCreate bool,category string,keyword string,rank_favorite,rank_time bool,collection_id int) ([]Item,error){
	var items []Item
	offset := (page_num -1) * page_size

	Condition := "where 1=1 "

	if rank_time == true {
		Condition = queryconditions + Condition + "order by created_at desc limit ? offset ?;"

		err := db.Select(&items, Condition,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Item{}, err
		}
		return items,nil
	}

	if rank_favorite == true {
		Condition = queryconditions + Condition + "order by like_count desc limit ? offset ?;"
		err := db.Select(&items, Condition,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Item{}, err
		}
		return items,nil
	}

	if len(category) > 0 {
		Condition = queryconditions + Condition + "and category = ? limit ? offset ?;"

		err := db.Select(&items, Condition,category,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Item{}, err
		}
		return items,nil
	}

	if len(keyword) > 0 {
		Condition = queryconditions +"where item_name like concat ('%',?,'%') limit ? offset ?;"
		err := db.Select(&items, Condition,keyword,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Item{}, err
		}
		return items,nil
	}
	fmt.Println("执行")
	if len(strconv.Itoa(collection_id)) >0 {
		Condition = queryconditions + Condition + "and collection_id = ? limit ? offset ?;"

		err := db.Select(&items, Condition,collection_id,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Item{}, err
		}
	}

	err := db.Select(&items, queryByCondition,page_size,offset)
	if err != nil {
		log.Println(err)
		return []Item{}, err
	}

	return items,nil
}

func GetCollectionList(page_num,page_size int64,userId string,keyword string,rank_favorite,rank_time bool,label string) ([]Collection,error){
	var collections []Collection
	offset := (page_num -1) * page_size

	Condition := "where 1=1 "

	if rank_time == true {
		Condition = collectionByCondition + Condition + "order by created_at desc limit ? offset ?;"

		err := db.Select(&collections, Condition,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Collection{}, err
		}
		return collections,nil
	}

	//if rank_favorite == true {
	//	Condition = collectionByCondition + Condition + "order by like_count desc limit ? offset ?;"
	//	err := db.Select(&collections, queryCollections,page_size,offset)
	//	if err != nil {
	//		log.Println(err)
	//		return []Collection{}, err
	//	}
	//	return collections,nil
	//}

	if len(keyword) > 0 {
		Condition = collectionByCondition +"where collection_name like concat ('%',?,'%') limit ? offset ?;"
		err := db.Select(&collections, Condition,keyword,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Collection{}, err
		}
		return collections,nil
	}


	err := db.Select(&collections, queryCollections,page_size,offset)
	if err != nil {
		log.Println(err)
		return []Collection{}, err
	}

	return collections,nil


}
