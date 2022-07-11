package model

import (
	"fmt"
	"log"
)

type ItemHistory struct {
	ItemID    string `db:"item_id"`
	From      string `db:"from"`
	To        string `db:"to"`
	Operation string `db:"operation"`
	Time      string `db:"time"`
}

var (
	addHistory   = string("inset into item_history(item_id,from,to,operation) values(?,?,?,?)")
	queryHistory = string("select * from item_history where item_id=?")
)

func AddMintHistory(itemId string, to string) error {
	from := ""
	operation := "Mint"
	_, e := db.Exec(addHistory, itemId, from, to, operation)
	if e != nil {
		return e
	}
	return nil
}

func AddTransferHistory(itemId string, from string, to string) error {
	operation := "Transfer"
	_, e := db.Exec(addHistory, itemId, from, to, operation)
	if e != nil {
		return e
	}
	return nil
}

func GetItemHistory(item_id string) ([]ItemHistory, error) {
	var ItemHistorys []ItemHistory

	err := db.Select(&ItemHistorys, queryHistory, item_id)
	if err != nil {
		log.Println(err)
		return []ItemHistory{}, err
	}

	fmt.Println(len(ItemHistorys))
	fmt.Println(ItemHistorys)
	return ItemHistorys, nil
}
