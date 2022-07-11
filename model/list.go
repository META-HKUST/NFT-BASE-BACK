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
)

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
		fmt.Println(Condition)
		err := db.Select(&items, Condition,category,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Item{}, err
		}
		return items,nil
	}

	if len(strconv.Itoa(collection_id)) >0 {
		Condition = queryconditions + Condition + "and collection_id = ? limit ? offset ?;"
		fmt.Println(Condition)
		err := db.Select(&items, Condition,collection_id,page_size,offset)
		if err != nil {
			log.Println(err)
			return []Item{}, err
		}
	}

	fmt.Println(len(items))
	fmt.Println(items)
	return items,nil
}
