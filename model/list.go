package model

import (
	"fmt"
	"log"
)

var (

	queryByCondition = string("select * from items limit ? offset ?")
	queryconditions = string("select * from items where ( category = ? or collection_id = ?) limit ? offset ?")
	queryByRankFavourite = string("select * from items order by rank_favourite desc limit ? offset ?")
	queryByRankTime = string("select * from items order by rank_time desc limit ? offset ?")
)

func GetItemList(page_num,page_size int64,category string,rank_favorite,rank_time bool,collection_id int) ([]Item,error){
	var items []Item
	offset := (page_num -1) * page_size
	if rank_time == true {
		err := db.Select(&items, queryByRankTime,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Item{}, err
		}
		return items,nil
	}
	if rank_favorite == true {
		err := db.Select(&items, queryByRankFavourite,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Item{}, err
		}
		return items,nil
	}

	err := db.Select(&items, queryconditions,category,collection_id,page_size,offset)
	if err != nil {
		log.Println(err)
		return []Item{}, err
	}

	fmt.Println(len(items))
	fmt.Println(items)
	return items,nil
}
