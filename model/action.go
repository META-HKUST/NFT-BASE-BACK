package model

import (
	"NFT-BASE-BACK/base"
	"errors"
	"fmt"
	"log"
)

var (
	//edit action
	addAction    = string("insert into action(act_name,creater_id,start_time,end_time,act_image,description,item_num) values(?,?,?,?,?,?,?);")
	deleteAction = string("delete from action where act_id = ?")
	getAction    = string("select * from action where act_id=?;")
	editAction   = string("update action set act_name=?,start_time=?,end_time=?,act_image=?,description=? where act_id=?;")
	uploadNFT    = string("insert into action_items(item_id, item_name, collection_id, item_data, description, owner_id, creater_id, category,like_count, act_id, vote_count) values(?,?,?,?,?,?,?,?,?,?,?);")
	// getItemList    = string()
	// vote           = string()
	getActionCount = string("select count(*) from action")
	getAllAct      = string("select * from action")
	getMaxActionId = string("select max(act_id) from action")

	getActionItemList = string("select item_id from action_item,action where action.act_id = action_item.act_id and action.act_id = ?;")
	queActItems       = string("select * from action_items where 1=1 and act_id = ?")
	queCanUpload      = string("select * from items where 1=1 and creater_id = ?")
	queItemIds        = string("select item_id from action_items where creater_id = ? and act_id = ?")
)

func GetCanUpload(page_num, page_size int64, act_id int64, userId string) ([]ItemAndLogo, error) {

	// initiate
	var items []Item
	var ItemAndLogos []ItemAndLogo
	offset := (page_num - 1) * page_size

	Condition := queCanUpload

	// construct query condition
	// select item_id that the user could not upload
	var ItemIds []string
	err1 := db.Select(&ItemIds, queItemIds, userId, act_id)
	if err1 != nil {
		log.Println(err1)
		return []ItemAndLogo{}, err1
	}
	log.Println("the id that user participated", ItemIds)

	for i := 0; i < len(ItemIds); i++ {
		Condition = Condition + " and item_id != '" + ItemIds[i] + "'"
	}

	Condition = Condition + " limit ? offset ?;"

	log.Println("can upload action items invoke sql: ", Condition)

	// select items that a user could upload
	err := db.Select(&items, Condition, userId, page_size, offset)
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
		ig.CoName, _ = GetCollectionName(items[i].CollectionID)
		ItemAndLogos = append(ItemAndLogos, ig)
	}

	return ItemAndLogos, nil
}

type Action struct {
	Act_id      int    `json:"act_id" db:"act_id"`
	Act_name    string `json:"act_name" db:"act_name"`
	Creater_id  string `json:"creater_id" db:"creater_id"`
	Create_time string `json:"create_time" db:"create_time"`
	Start_time  string `json:"start_time" db:"start_time"`
	End_time    string `json:"end_time" db:"end_time"`
	Act_image   string `json:"act_image" db:"act_image"`
	Description string `json:"description" db:"description"`
	Item_num    int    `json:"item_num" db:"item_num"`
}

// this struction represents a single item details
type ActItem struct {
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
	ActID        int    `json:"act_id" db:"act_id"`
	VoteCount    int    `json:"vote_count" db:"vote_count"`
}

type ActAndVote struct {
	ActItem
	Vote      bool   `json:"vote" db:"vote"`
	LogoImage string `json:"logo_image" db:"logo_image"`
	CoName    string `json:"collection_name" db:"collection_name"`
}

func GetActItemList(page_num, page_size int64, act_id int64, rank_vote bool, rank_time bool, userId string, keyword string) ([]ActAndVote, error) {
	var actItems []ActItem
	var ItemAndLikes []ActAndVote
	offset := (page_num - 1) * page_size

	if keyword != "" {
		Condition := queActItems + " and item_name like concat ('%',?,'%') "
		if rank_time == true {
			Condition = Condition + " order by created_at desc"
		} else if rank_vote == true {
			Condition = Condition + " order by vote_count desc"
		}
		Condition = Condition + " limit ? offset ?;"

		log.Println("query action items condition: ", Condition)
		err := db.Select(&actItems, Condition, act_id, keyword, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ActAndVote{}, err
		}
		// add like and logoImage
		for i := 0; i < len(actItems); i++ {
			ig := ActAndVote{}
			ig.ActItem = actItems[i]
			ig.Vote, _ = DoesVote(actItems[i].ActID, actItems[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(actItems[i].CollectionID)
			ig.LogoImage, _ = GetLogoImage(actItems[i].OwnerID)
			ItemAndLikes = append(ItemAndLikes, ig)
		}
		return ItemAndLikes, nil

	}
	if rank_vote == true {
		Condition := queActItems + " order by vote_count desc limit ? offset ?;"
		log.Println("query action items condition: ", Condition)
		err := db.Select(&actItems, Condition, act_id, page_size, offset)

		if err != nil {
			log.Println(err)
			return []ActAndVote{}, err
		}
		// add like and logoImage
		for i := 0; i < len(actItems); i++ {
			ig := ActAndVote{}
			ig.ActItem = actItems[i]
			ig.Vote, _ = DoesVote(actItems[i].ActID, actItems[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(actItems[i].CollectionID)
			ig.LogoImage, _ = GetLogoImage(actItems[i].OwnerID)
			ItemAndLikes = append(ItemAndLikes, ig)
		}
		return ItemAndLikes, nil
	}
	if rank_time == true {
		Condition := queActItems + " order by created_at desc limit ? offset ?;"
		log.Println("query action items condition: ", Condition)
		err := db.Select(&actItems, Condition, act_id, page_size, offset)
		if err != nil {
			log.Println(err)
			return []ActAndVote{}, err
		}
		// add like and logoImage
		for i := 0; i < len(actItems); i++ {
			ig := ActAndVote{}
			ig.ActItem = actItems[i]
			ig.Vote, _ = DoesVote(actItems[i].ActID, actItems[i].ItemID, userId)
			ig.CoName, _ = GetCollectionName(actItems[i].CollectionID)
			ig.LogoImage, _ = GetLogoImage(actItems[i].OwnerID)
			ItemAndLikes = append(ItemAndLikes, ig)
		}
		return ItemAndLikes, nil

	}

	Condition := queActItems + " limit ? offset ?;"
	err := db.Select(&actItems, Condition, act_id, page_size, offset)
	log.Println("query action items condition: ", Condition)
	if err != nil {
		log.Println(err)
		return []ActAndVote{}, err
	}
	// add like and logoImage
	for i := 0; i < len(actItems); i++ {
		ig := ActAndVote{}
		ig.ActItem = actItems[i]
		ig.Vote, _ = DoesVote(actItems[i].ActID, actItems[i].ItemID, userId)
		ig.CoName, _ = GetCollectionName(actItems[i].CollectionID)
		ig.LogoImage, _ = GetLogoImage(actItems[i].OwnerID)
		ItemAndLikes = append(ItemAndLikes, ig)
	}
	return ItemAndLikes, nil

}

type ActionItem struct {
	ItemInfo
	VoteCount int `json:"vote_count" db:"vote_count"`
}

func GetActionItemList(act_id int, rank_vote bool, rank_time bool) ([]ActionItem, error) {

	var ActionItems []ActionItem
	var itemIds []string

	err := db.Select(&itemIds, getActionItemList, act_id)
	if err != nil {
		log.Println(err)
		return ActionItems, err
	}

	log.Println("item ids: ", itemIds)

	for i := 0; i < len(itemIds); i++ {
		ai := ActionItem{}
		ai.ItemInfo, _ = GetItemInfo(itemIds[i])
		ai.VoteCount, _ = GetVoteCount(act_id, itemIds[i])
		ActionItems = append(ActionItems, ai)
	}

	return ActionItems, nil
}

func GetMaxActionId() (int, error) {
	var g int
	err := db.Get(&g, getMaxActionId)
	if err != nil {
		return -1, err
	}
	return g, nil
}

func AddAction(act_name string, creater_id string, start_time string, end_time string, act_image string, description string, item_num int) (Action, error) {

	_, err := db.Exec(addAction, act_name, creater_id, start_time, end_time, act_image, description, item_num)
	if err != nil {
		log.Println(err)
		return Action{}, err
	}

	max, _ := GetMaxActionId()
	a, e := GetAction(max)
	if e != nil {
		log.Println("Transaction start failed", e)
		return Action{}, e
	}
	return a, nil
}

func DeleteAction(act_id int) error {
	_, e := db.Exec(deleteAction, act_id)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return e
	}
	return nil
}

func EditAction(act_id int, act_name string, start_time string, end_time string, act_image string, description string) (Action, error) {

	a, e := GetAction(act_id)

	if act_name != "" {
		a.Act_name = act_name
	}
	if start_time != "" {
		a.Start_time = start_time
	}
	if end_time != "" {
		a.End_time = end_time
	}
	if act_image != "" {
		a.Act_image = act_image
	}
	if description != "" {
		a.Description = description
	}
	result, e := db.Exec(editAction, a.Act_name, a.Start_time, a.End_time, a.Act_image, a.Description, a.Act_id)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return Action{}, e
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	log.Println("rowsAffected: ", rowsAffected, "lastInsertId: ", lastInsertId)

	a, e = GetAction(act_id)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return Action{}, e
	}
	return a, nil
}

func GetAction(act_id interface{}) (Action, error) {
	var a Action
	fmt.Println("act_id: ", act_id)
	err := db.Get(&a, getAction, act_id)
	if err != nil {
		log.Println(err)
		return Action{}, err
	}
	return a, nil
}

func UploadNFT(act_id int, item_id string) error {

	act, _ := GetAction(act_id)
	var act1 Action
	item, _ := GetItemInfo(item_id)
	if act == act1 {
		return errors.New("can not find this action in database")
	}

	_, e := db.Exec(uploadNFT, item.ItemID, item.ItemName, item.CollectionID, item.ItemData, item.Description, item.OwnerID, item.CreaterID, item.Category, item.LikeCount, act_id, 0)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return e
	}
	return nil
}

func GetActionCount() (int, error) {
	var a int
	err := db.Get(&a, getActionCount)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return a, nil
}

func GetAllAct() ([]Action, error) {
	var Actions []Action

	err := db.Select(&Actions, getAllAct)
	if err != nil {
		log.Println(err)
		return Actions, err
	}

	return Actions, nil
}

func ItemList() {

}
