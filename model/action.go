package model

import (
	"NFT-BASE-BACK/base"
	"fmt"
	"log"
)

var (
	//edit user profile
	addAction    = string("insert into action(act_name,creater_id,start_time,end_time,act_image,description,item_num) values(?,?,?,?,?,?,?);")
	deleteAction = string("delete from action where act_id = ?")
	getAction    = string("select * from action where act_id=?;")
	editAction   = string("update action set act_name=?,start_time=?,end_time=?,act_image=?,description=? where act_id=?;")
	uploadNFT    = string("insert into action_item(act_id,item_id) values(?,?);")
	// getItemList    = string()
	// vote           = string()
	getActionCount = string("select count(*) from action")
	getAllAct      = string("select * from action")
	getMaxActionId = string("select max(act_id) from action")
)

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
	_, e := db.Exec(uploadNFT, act_id, item_id)
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
		fmt.Println(len(Actions))
		if len(Actions) == 0 {
			Actions = append(Actions, Action{
				1,
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				0,
			})
		}
		log.Println(err)
		return Actions, err
	}
	log.Println(len(Actions))
	if len(Actions) == 0 {
		Actions = append(Actions, Action{
			1,
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			0,
		})
	}
	return Actions, nil
}

func ItemList() {

}
