package model

import (
	"log"
	"strconv"
)

var (
	queryByCondition      = string("select * from items limit ? offset ?")
	queryconditions       = string("select * from items ")
	queryByRankFavourite  = string("select * from items order by rank_favourite desc limit ? offset ?")
	queryByRankTime       = string("select * from items order by rank_time desc limit ? offset ?")
	queryCollections      = string("select * from collection limit ? offset ?")
	collectionByCondition = string("select * from collection ")
)

type Collection struct {
	CollectionId   string `json:"collection_id" db:"collection_id"`
	CollectionName string `json:"collection_name" db:"collection_name"`
	BannerImage    string `json:"banner_image" db:"banner_image"`
	LogoImage      string `json:"logo_image" db:"logo_image"`
	FeatureImage   string `json:"feature_image" db:"feature_image"`
	Description    string `json:"description" db:"description"`
	ItemNum        int    `json:"item_num" db:"items_count"`
	OwnerId        string `json:"owner_id" db:"owner"`
	OwnerName      string `json:"owner_name" db:"owner_name"`
	CreationTime   string `json:"creation_time" db:"created_at"`
}

type ItemAndLogo struct {
	Item
	LogoImage string `json:"logo_image" db:"logo_image"`
	Like      bool   `json:"like" db:"like"`
	CoName    string `json:"collection_name" db:"collection_name"`
}

func GetItemList(page_num, page_size int64, userId string, userLike, userCollect, userCreate bool, category string, keyword string, rank_favorite, rank_time bool, collection_id int) ([]ItemAndLogo, error) {
	var items []Item
	var ItemAndLogos []ItemAndLogo
	offset := (page_num - 1) * page_size

	Condition := "where 1=1 "

	if rank_time == true {
		Condition = queryconditions + Condition + "order by created_at desc limit ? offset ?;"

		err := db.Select(&items, Condition, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, err
		}
		// add like and logoImage
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(ig.CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		return ItemAndLogos, nil
	}

	if userCollect == true {
		Condition = queryconditions + Condition + "and owner_id = ? limit ? offset ?;"

		err := db.Select(&items, Condition, category, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, err
		}
		// add like and logoImage
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(ig.CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		return ItemAndLogos, nil
	}

	if rank_favorite == true {
		Condition = queryconditions + Condition + "order by like_count desc limit ? offset ?;"
		err := db.Select(&items, Condition, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, err
		}
		// add like and logoImage
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(ig.CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		return ItemAndLogos, nil
	}

	if len(category) > 0 {
		Condition = queryconditions + Condition + "and category = ? limit ? offset ?;"

		err := db.Select(&items, Condition, category, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, err
		}
		// add like and logoImage
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(ig.CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		return ItemAndLogos, nil
	}

	if len(keyword) > 0 {
		Condition = queryconditions + "where item_name like concat ('%',?,'%') limit ? offset ?;"
		err := db.Select(&items, Condition, keyword, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, err
		}
		// add like and logoImage
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(ig.CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		return ItemAndLogos, nil
	}

	if len(strconv.Itoa(collection_id)) > 0 {
		Condition = queryconditions + Condition + "and collection_id = ? limit ? offset ?;"

		err := db.Select(&items, Condition, collection_id, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, err
		}
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.CoName, _ = GetCollectionName(ig.CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		return ItemAndLogos, nil
	}

	err := db.Select(&items, queryByCondition, page_size, offset)
	if err != nil {
		log.Println(err)
		return []ItemAndLogo{}, err
	}

	for i := 0; i < len(items); i++ {
		ig := ItemAndLogo{}
		ig.Item = items[i]
		ig.Like, _ = DoesLike(items[i].ItemID, userId)
		ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
		ig.CoName, _ = GetCollectionName(ig.CollectionID)
		ItemAndLogos = append(ItemAndLogos, ig)
	}
	return ItemAndLogos, nil
}

func GetCollectionList(page_num, page_size int64, userId string, keyword string, rank_favorite, rank_time bool, label string) ([]Collection, error) {
	var collections []Collection
	offset := (page_num - 1) * page_size

	Condition := "where 1=1 "

	if rank_time == true {
		Condition = collectionByCondition + Condition + "order by created_at desc limit ? offset ?;"

		err := db.Select(&collections, Condition, page_size, offset)
		if err != nil {
			log.Println(err)
			return []Collection{}, err
		}
		return collections, nil
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
		Condition = collectionByCondition + "where collection_name like concat ('%',?,'%') limit ? offset ?;"
		err := db.Select(&collections, Condition, keyword, page_size, offset)
		if err != nil {
			log.Println(err)
			return []Collection{}, err
		}
		return collections, nil
	}

	if len(label) > 0 {
		Condition = collectionByCondition + "where collection_id in ( select collection_id from collection_label where label like concat ('%',?,'%')) limit ? offset ?;"
		err := db.Select(&collections, Condition, label, page_size, offset)
		if err != nil {
			log.Println(err)
			return []Collection{}, err
		}
		return collections, nil
	}

	err := db.Select(&collections, queryCollections, page_size, offset)
	if err != nil {
		log.Println(err)
		return []Collection{}, err
	}

	return collections, nil

}
