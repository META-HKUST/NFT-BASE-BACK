package model

import (
	"errors"
	"log"
)

var (
	addVote         = string("insert into item_vote(act_id,item_id,user_id) values(?,?,?)")
	deleteVote      = string("delete from item_vote where act_id=? and item_id = ? and user_id=?")
	getVoteCount    = string("select vote_count from action_items where act_id=? and item_id = ?")
	doesVote        = string("select item_id from item_vote where act_id=? and item_id = ? and user_id = ?")
	updateVoteCount = string("update action_items set vote_count = ? where act_id = ? and item_id = ?")
)

func Vote(actId int, itemId string, UserId string) error {
	act, _ := GetAction(actId)
	var act1 Action
	if act == act1 {
		return errors.New("can not find this action in database")
	}
	_, e := db.Exec(addVote, actId, itemId, UserId)
	if e != nil {
		log.Println(e)
		return e
	}
	// update vote_count in database
	var count int
	_ = db.Get(&count, getVoteCount, actId, itemId)
	_, e = db.Exec(updateVoteCount, count+1, actId, itemId)
	if e != nil {
		log.Println(e)
		return e
	}

	return nil
}

func UnVote(actId int, itemId string, UserId string) error {

	act, _ := GetAction(actId)
	var act1 Action
	if act == act1 {
		return errors.New("can not find this action in database")
	}
	_, e := db.Exec(deleteVote, actId, itemId, UserId)
	if e != nil {
		log.Println(e)
		return e
	}

	// update vote_count in database
	var count int
	_ = db.Get(&count, getVoteCount, actId, itemId)
	_, e = db.Exec(updateVoteCount, count-1, actId, itemId)
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}

func DoesVote(actId int, itemId string, UserId string) (bool, error) {
	var g string
	e := db.Get(&g, doesVote, actId, itemId, UserId)
	if e != nil {
		return false, e
	}
	if g == "" {
		return false, nil
	} else {
		return true, nil
	}
}

func GetVoteCount(actId int, itemId string) (int, error) {
	var a int
	err := db.Get(&a, getVoteCount, actId, itemId)
	if err != nil {
		return 0, err
	}
	return a, nil
}
