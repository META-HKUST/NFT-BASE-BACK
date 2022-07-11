package model

import (
	"NFT-BASE-BACK/base"
	"log"
)

var (
	//edit user profile
	addAction    = string("insert into action(act_name,creater_id,start_time,end_time,act_image,description,item_num) values(?,?,?,?,?,?,?,?);")
	deleteAction = string("delete from action where act_id = ?")
	getAction    = string("select * from action where act_id=?;")
	editAction   = string("update action set act_name=?,start_time=?,end_time=?,act_image=?,description=? where act_id=?;")
	uploadNFT    = string("insert into action_item(act_id,item_id) values(?,?);")
	// getItemList    = string()
	// vote           = string()
	getActionCount = string("select count(*) from action")

	getMaxActionId = string("select max(collection_id) from action")
)

type Action struct {
	act_id      int    `json:"act_id" db:"act_id"`
	act_name    string `json:"act_name" db:"act_name"`
	creater_id  string `json:"creater_id" db:"creater_id"`
	create_time string `json:"create_time" db:"create_time"`
	start_time  string `json:"start_time" db:"start_time"`
	end_time    string `json:"end_time" db:"end_time"`
	act_image   string `json:"act_image" db:"act_image"`
	description string `json:"description" db:"description"`
	item_num    int    `json:"item_num" db:"item_num"`
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
	result, e := db.Exec(addAction, act_name, creater_id, start_time, end_time, act_image, description, item_num)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return Action{}, e
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	log.Println("rowsAffected: ", rowsAffected, "lastInsertId: ", lastInsertId)

	max, _ := GetMaxCollectionId()
	a, _ := GetAction(max)
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
		a.act_name = act_name
	}
	if start_time != "" {
		a.start_time = start_time
	}
	if end_time != "" {
		a.end_time = end_time
	}
	if act_image != "" {
		a.act_image = act_image
	}
	if description != "" {
		a.description = description
	}
	result, e := db.Exec(editAction, a.act_name, a.start_time, a.end_time, a.act_image, a.description, a.act_id)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return Action{}, e
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	log.Println("rowsAffected: ", rowsAffected, "lastInsertId: ", lastInsertId)

	return a, nil
}

func GetAction(act_id int) (Action, error) {
	var a Action
	err := db.Get(&a, getAction, act_id)
	if err != nil {
		return Action{}, err
	}
	return a, nil
}

func UploadNFT(act_id int, item_id string) error {
	_, e := db.Exec(uploadNFT, act_id, item_id)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return e
	}
	return nil
}

func ItemList() {

}
