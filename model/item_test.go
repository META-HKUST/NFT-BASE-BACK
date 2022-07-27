package model

import (
	"NFT-BASE-BACK/config"
	"log"
	"testing"
)

func TestCreateItem(t *testing.T) {
	err := config.LoadConfig("../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	InitDB(config.CONFIG)

	item := Item{
		ItemID:       "1",
		ItemName:     "jahha",
		CollectionID: "1",
		ItemData:     "ssfa",
		Description:  "1111",
		OwnerID:      "sdswe",
		CreaterID:    "wef",
		Category:     "sdfe",
	}

	ret, err := CreateItem(item)
	if err != nil {
		log.Println(err)
	}
	log.Println(ret)
}

func TestCreateItemLabel(t *testing.T) {
	err := config.LoadConfig("../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	InitDB(config.CONFIG)

	itemLable := ItemLable{
		ItemID:    "1",
		ItemLabel: "hah",
	}

	ret, err := CreateItemLabel(itemLable)
	if err != nil {
		log.Println(err)
	}
	log.Println(ret)
}

func TestSearchLabel(t *testing.T) {
	err := config.LoadConfig("../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	InitDB(config.CONFIG)

	itemLable := ItemLable{
		ItemID:    "1",
		ItemLabel: "hah",
	}

	_, err = CreateItemLabel(itemLable)
	if err != nil {
		log.Println(err)
	}

	itemLable = ItemLable{
		ItemID:    "1",
		ItemLabel: "hahaaa",
	}

	_, err = CreateItemLabel(itemLable)
	if err != nil {
		log.Println(err)
	}

	ret2, err := SearchLable("1")

	if err != nil {
		log.Println(err)
	}
	log.Println(ret2)
}
