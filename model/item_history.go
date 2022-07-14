package model

import (
	"log"
)

type ItemHistory struct {
	ItemID    string `db:"item_id"`
	From      string `db:"from_id"`
	To        string `db:"to_id"`
	Operation string `db:"operation"`
	Time      string `db:"time"`
}

var (
	addHistory   = string("insert into item_history(item_id,from_id,to_id,operation) values(?,?,?,?)")
	queryHistory = string("select * from item_history where item_id=?;")
)

func AddMintHistory(itemId string, to string) error {

	from := "Admin"
	operation := "Mint"
	_, e := db.Exec(addHistory, itemId, from, to, operation)
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}

func AddTransferHistory(itemId string, from string, to string) error {
	operation := "Transfer"
	_, e := db.Exec(addHistory, itemId, from, to, operation)
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}

type HistoryRes struct {
	His      ItemHistory
	FromName string
	ToName   string
}

func GetItemHistory(item_id string) ([]HistoryRes, error) {
	var ItemHistorys []ItemHistory

	err := db.Select(&ItemHistorys, queryHistory, item_id)
	if err != nil {
		log.Println(err)
		return []HistoryRes{}, err
	}

	var hisres []HistoryRes

	for _, v := range ItemHistorys {
		from, e1 := GetUserName(v.From)
		if e1 != nil {
			log.Println(e1)
		}
		to, e2 := GetUserName(v.From)
		if e2 != nil {
			log.Println(e2)
		}
		hisres = append(hisres, HistoryRes{
			v,
			from,
			to,
		})
	}

	return hisres, nil
}
