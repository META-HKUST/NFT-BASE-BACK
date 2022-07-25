package model

import (
	"log"
)

var (
	queryByCondition      = string("select * from items limit ? offset ?")
	queryconditions       = string("select * from items ")
	queryByRankFavourite  = string("select * from items order by rank_favourite desc limit ? offset ?")
	queryByRankTime       = string("select * from items order by rank_time desc limit ? offset ?")
	queryCollections      = string("select * from collection limit ? offset ?")
	collectionByCondition = string("select * from collection ")
	queryCoCount          = string("select count(*) from collection")
	queryItemCount        = string("select count(*) from items")
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

func GetItemList(page_num, page_size int64, userId string, userLike, userCollect, userCreate bool, category string, keyword string, rank_favorite, rank_time bool, collection_id int) ([]ItemAndLogo, int, error) {
	var items []Item
	var ItemAndLogos []ItemAndLogo
	offset := (page_num - 1) * page_size

	Condition := "where 1=1 "

	if userCollect == true {

		// select first and then rank
		Condition = queryconditions + Condition + "and owner_id = ?  "

		if rank_time == true {
			Condition = Condition + " order by created_at desc"
		} else if rank_favorite == true {
			Condition = Condition + " order by like_count desc"
		}
		Condition = Condition + " limit ? offset ?;"

		err := db.Select(&items, Condition, userId, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, 0, err
		}
		// add like and logoImage
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(items[i].CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		// select count
		queryC := queryItemCount + " where owner_id = ?"
		var count []int
		err1 := db.Select(&count, queryC, userId)
		if err != nil {
			log.Println(err1)
			return []ItemAndLogo{}, 0, err1
		}

		return ItemAndLogos, count[0], nil
	}

	if userCreate == true {

		// select first and then rank
		Condition = queryconditions + Condition + "and creater_id = ?  "

		if rank_time == true {
			Condition = Condition + " order by created_at desc"
		} else if rank_favorite == true {
			Condition = Condition + " order by like_count desc"
		}
		Condition = Condition + " limit ? offset ?;"

		err := db.Select(&items, Condition, userId, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, 0, err
		}
		// add like and logoImage
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(items[i].CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		// select count
		queryC := queryItemCount + " where owner_id = ?"
		var count []int
		err1 := db.Select(&count, queryC, userId)
		if err != nil {
			log.Println(err1)
			return []ItemAndLogo{}, 0, err1
		}

		return ItemAndLogos, count[0], nil
	}

	if len(category) > 0 {

		// select first and then rank
		Condition = queryconditions + Condition + "and category = ?  "

		if rank_time == true {
			Condition = Condition + " order by created_at desc"
		} else if rank_favorite == true {
			Condition = Condition + " order by like_count desc"
		}
		Condition = Condition + " limit ? offset ?;"

		err := db.Select(&items, Condition, category, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, 0, err
		}
		// add like and logoImage
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(items[i].CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}

		if rank_favorite == true {
			Condition = queryconditions + Condition + "order by like_count desc limit ? offset ?;"
			err := db.Select(&items, Condition, page_size, offset)
			if err != nil {
				log.Println(err)
				return []ItemAndLogo{}, 0, err
			}
			// add like and logoImage
			for i := 0; i < len(items); i++ {
				ig := ItemAndLogo{}
				ig.Item = items[i]
				ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
				ig.Like, _ = DoesLike(items[i].ItemID, userId)
				ig.CoName, _ = GetCollectionName(items[i].CollectionID)
				ItemAndLogos = append(ItemAndLogos, ig)
			}
			// select count
			var count []int
			err1 := db.Select(&count, queryItemCount)
			if err != nil {
				log.Println(err1)
			}

			return ItemAndLogos, count[0], nil
		}

		if rank_time == true {
			Condition = queryconditions + Condition + "order by created_at desc limit ? offset ?;"

			err := db.Select(&items, Condition, page_size, offset)
			if err != nil {
				log.Println(err)
				return []ItemAndLogo{}, 0, err
			}
			// add like and logoImage
			for i := 0; i < len(items); i++ {
				ig := ItemAndLogo{}
				ig.Item = items[i]
				ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
				ig.Like, _ = DoesLike(items[i].ItemID, userId)
				ig.CoName, _ = GetCollectionName(items[i].CollectionID)
				ItemAndLogos = append(ItemAndLogos, ig)
			}
			// select count
			var count []int
			err1 := db.Select(&count, queryItemCount)
			if err != nil {
				log.Println(err1)
			}

			return ItemAndLogos, count[0], nil
		}

		// select count
		queryC := queryItemCount + " where category = ?"
		var count []int
		err1 := db.Select(&count, queryC, category)
		if err != nil {
			log.Println(err1)
		}

		return ItemAndLogos, count[0], nil
	}

	if len(keyword) > 0 {
		Condition = queryconditions + "where item_name like concat ('%',?,'%') limit ? offset ?;"
		err := db.Select(&items, Condition, keyword, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, 0, err
		}
		// add like and logoImage
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.Like, _ = DoesLike(items[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(items[i].CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		// select count
		queryC := queryItemCount + " where item_name like concat ('%',?,'%')"
		var count []int
		err1 := db.Select(&count, queryC, keyword)
		if err != nil {
			log.Println(err1)
		}

		return ItemAndLogos, count[0], nil
	}

	if collection_id > 0 {
		Condition = queryconditions + Condition + "and collection_id = ? limit ? offset ?;"

		err := db.Select(&items, Condition, collection_id, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ItemAndLogo{}, 0, err
		}
		for i := 0; i < len(items); i++ {
			ig := ItemAndLogo{}
			ig.Item = items[i]
			ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
			ig.CoName, _ = GetCollectionName(items[i].CollectionID)
			ItemAndLogos = append(ItemAndLogos, ig)
		}
		// select count
		queryC := queryItemCount + " where collection_id = ?"
		var count []int
		err1 := db.Select(&count, queryC, collection_id)
		if err != nil {
			log.Println(err1)
			return []ItemAndLogo{}, 0, err1
		}

		return ItemAndLogos, count[0], nil
	}

	err := db.Select(&items, queryByCondition, page_size, offset)
	if err != nil {
		log.Println(err)
		return []ItemAndLogo{}, 0, err
	}

	for i := 0; i < len(items); i++ {
		ig := ItemAndLogo{}
		ig.Item = items[i]
		ig.LogoImage, _ = GetLogoImage(items[i].CreaterID)
		ig.CoName, _ = GetCollectionName(items[i].CollectionID)
		ItemAndLogos = append(ItemAndLogos, ig)
	}
	// select count
	var count []int
	err1 := db.Select(&count, queryItemCount)
	if err != nil {
		log.Println(err1)
	}

	return ItemAndLogos, count[0], nil
}

func GetCollectionList(page_num, page_size int64, userId string, keyword string, rank_favorite, rank_time bool, label string) ([]Collection, int, error) {
	var collections []Collection
	offset := (page_num - 1) * page_size

	Condition := "where 1=1 "

	if rank_time == true {
		Condition = collectionByCondition + Condition + "order by created_at desc limit ? offset ?;"

		err := db.Select(&collections, Condition, page_size, offset)
		if err != nil {
			log.Println(err)
			return []Collection{}, 0, err
		}
		// select count
		queryC := queryCoCount
		var count []int
		err1 := db.Select(&count, queryC)
		if err != nil {
			log.Println(err1)
		}
		return collections, count[0], nil
	}

	if len([]byte(userId)) > 0 {
		log.Println(userId)
		Condition = collectionByCondition + Condition + "and owner=? limit ? offset ?;"

		err := db.Select(&collections, Condition, userId, page_size, offset)
		if err != nil {
			log.Println(err)
			return []Collection{}, 0, err
		}
		// select count
		queryC := queryCoCount + " where owner=?"
		var count []int
		err1 := db.Select(&count, queryC, userId)
		if err != nil {
			log.Println(err1)
		}
		return collections, count[0], nil
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
			return []Collection{}, 0, err
		}
		// select count
		queryC := queryCoCount + " where collection_name like concat ('%',?,'%')"
		var count []int
		err1 := db.Select(&count, queryC, keyword)
		if err != nil {
			log.Println(err1)
		}
		return collections, count[0], nil
	}

	if len(label) > 0 {
		Condition = collectionByCondition + "where collection_id in ( select collection_id from collection_label where label like concat ('%',?,'%')) limit ? offset ?;"
		err := db.Select(&collections, Condition, label, page_size, offset)
		if err != nil {
			log.Println(err)
			return []Collection{}, 0, err
		}
		// select count
		queryC := queryCoCount + " where collection_id in ( select collection_id from collection_label where label like concat ('%',?,'%'))"
		var count []int
		err1 := db.Select(&count, queryC, label)
		if err != nil {
			log.Println(err1)
		}
		return collections, count[0], nil
	}

	err := db.Select(&collections, queryCollections, page_size, offset)
	if err != nil {
		log.Println(err)
		return []Collection{}, 0, err
	}
	// select count
	queryC := queryCoCount
	var count []int
	err1 := db.Select(&count, queryC)
	if err != nil {
		log.Println(err1)
		return []Collection{}, 0, err1
	}
	return collections, count[0], nil
}
