package model

import (
	"errors"
	"log"
)

const (
	insertItem      = string("insert into items(item_id, item_name, collection_id, item_data, description, owner_id, creater_id, category,like_count) values(?,?,?,?,?,?,?,?,?);")
	queryItem       = string("select item_id, item_name, collection_id, item_data, description, owner_id, creater_id, category, created_at from items where item_id=?;")
	insertItemLable = string("insert into item_label(item_id, label) values(?,?);")
	queryItemLable  = string("select item_id, label from item_label where item_id=?")
	updateItemOwner = string("update items set owner_id=? where item_id=?;")
	searchLabel     = string("select label from item_label where item_id=?;")
	updateItemLabel = string("update item_label set label=? where item_id=?;")
	QueryItem       = string("select * from items where item_id=?;")
)

type Item struct {
	ItemID       string `json:"item_id" db:"item_id"`
	ItemName     string `json:"item_name" db:"item_name"`
	CollectionID string `json:"collection_id" db:"collection_id"`
	ItemData     string `json:"item_data" db:"item_data"`
	Description  string `json:"description" db:"description"`
	OwnerID      string `json:"owner_id" db:"owner_id"`
	CreaterID    string `json:"creater_id" db:"creater_id"`
	Category     string `json:"category" db:"category"`
	LikeCount    int    `json:"like_count" db:"like_count"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}

// add label in except to Item
type ItemInfo struct {
	ItemName     string   `json:"item_name" db:"item_name"`
	ItemID       string   `json:"item_id" db:"item_id"`
	ItemData     string   `json:"item_data" db:"item_data"`
	CreatedTime  string   `json:"created_time" db:"created_at"`
	Description  string   `json:"description" db:"description"`
	CollectionID string   `json:"collection_id" db:"collection_id"`
	Category     string   `json:"category" db:"category"`
	Label        []string `json:"label" db:"label"`
	CreaterID    string   `json:"creater_id" db:"creater_id"`
	OwnerID      string   `json:"owner_id" db:"owner_id"`
	LikeCount    int      `json:"like_count" db:"like_count"`
}

type ItemInfoAndLike struct {
	ItemInfo
	Like bool `json:"like" db:"like"`
}

func CreateItem(item Item) (Item, error) {
	_, err := db.Exec(insertItem,
		item.ItemID,
		item.ItemName,
		item.CollectionID,
		item.ItemData,
		item.Description,
		item.OwnerID,
		item.CreaterID,
		item.Category,
		0,
	)
	if err != nil {
		return Item{}, err
	}
	var ret Item
	err = db.Get(&ret, queryItem, item.ItemID)
	if err != nil {
		return Item{}, err
	}
	return ret, nil
}

type ItemLable struct {
	ItemID    string `db:"item_id"`
	ItemLabel string `db:"label"`
}

func CreateItemLabel(label ItemLable) (ItemLable, error) {
	_, err := db.Exec(insertItemLable,
		label.ItemID,
		label.ItemLabel,
	)
	if err != nil {
		return ItemLable{}, err
	}
	var ret ItemLable
	err = db.Get(&ret, queryItemLable, label.ItemID)
	if err != nil {
		return ItemLable{}, err
	}
	return ret, nil
}

func UpdateItemOwner(itemId, toUserId string) (Item, error) {
	_, err := db.Exec(updateItemOwner,
		toUserId,
		itemId,
	)
	if err != nil {
		return Item{}, err
	}
	var ret Item
	err = db.Get(&ret, queryItem, itemId)
	if err != nil {
		return Item{}, err
	}
	return ret, nil
}

type Label struct {
	Label string `db:"label"`
}
type LabelSlice struct {
	Label []Label
}

func SearchLable(itemID string) ([]string, error) {
	var LableSlice []string
	err := db.Select(&LableSlice, searchLabel, itemID)
	if err != nil {
		return []string{}, err
	}
	return LableSlice, nil
}

func EditItemLable(label []string, itemID string) error {

	labelss, _ := SearchLable(itemID)
	for i := 0; i < len(label); i++ {
		contain := false
		for j := 0; j < len(labelss); j++ {
			if labelss[j] == label[i] {
				contain = true
			}
		}
		if contain == false {
			_, err := db.Exec(insertItemLable,
				itemID,
				label[i],
			)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}

func EditItem(itemId, itemName, description, collectionId string, label []string) error {

	if itemId == "" {
		return errors.New("Parameter input error")
	}

	argsItem := []string{}

	updateItemInfo := "update items set "
	if itemName != "" {
		updateItemInfo = updateItemInfo + "item_name=?,"
		argsItem = append(argsItem, itemName)
	}

	if description != "" {
		updateItemInfo = updateItemInfo + "description=?,"
		argsItem = append(argsItem, description)
	}
	if collectionId != "" {
		updateItemInfo = updateItemInfo + "collection_id=?,"
		argsItem = append(argsItem, collectionId)
	}

	updateItemInfo = updateItemInfo[:len(updateItemInfo)-1]
	updateItemInfo = updateItemInfo + " " + "where item_id=?;"
	argsItem = append(argsItem, itemId)

	paramItem := make([]interface{}, len(argsItem))
	for i, v := range argsItem {
		paramItem[i] = v
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Println("Transaction start failed", err)
		return err
	}

	tx.MustExec(updateItemInfo, paramItem...)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	err = EditItemLable(label, itemId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func GetItemInfo(itemId string) (ItemInfo, error) {
	var item Item
	var lableSlice []string
	err := db.Get(&item, QueryItem, itemId)
	if err != nil {
		log.Println(err)
		return ItemInfo{}, err
	}

	err = db.Select(&lableSlice, searchLabel, itemId)
	if err != nil {
		log.Println(err)
		return ItemInfo{}, err
	}

	itemInfo := ItemInfo{
		item.ItemName,
		item.ItemID,
		item.ItemData,
		item.CreatedAt,
		item.Description,
		item.CollectionID,
		item.Category,
		lableSlice,
		item.CreaterID,
		item.OwnerID,
		item.LikeCount,
	}

	return itemInfo, nil
}

func GetItemAndLikeInfo(itemId string, userId string) (ItemInfoAndLike, error) {
	var item Item
	var lableSlice []string
	err := db.Get(&item, QueryItem, itemId)
	if err != nil {
		log.Println(err)
		return ItemInfoAndLike{}, err
	}

	err = db.Select(&lableSlice, searchLabel, itemId)
	if err != nil {
		log.Println(err)
		return ItemInfoAndLike{}, err
	}

	itemInfo := ItemInfo{
		item.ItemName,
		item.ItemID,
		item.ItemData,
		item.CreatedAt,
		item.Description,
		item.CollectionID,
		item.Category,
		lableSlice,
		item.CreaterID,
		item.OwnerID,
		item.LikeCount,
	}

	itemAndLike := ItemInfoAndLike{}
	itemAndLike.ItemInfo = itemInfo
	itemAndLike.Like, _ = DoesLike(itemInfo.ItemID, userId)

	return itemAndLike, nil
}
