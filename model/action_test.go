package model

import (
	"NFT-BASE-BACK/config"
	"log"
	"testing"
)

func TestActionItemList(t *testing.T) {
	err := config.LoadConfig("../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	InitDB(config.CONFIG)

	ret, _ := GetActionItemList(1, false, false)

	log.Println(ret)
}

func TestGetVoteCount(t *testing.T) {
	err := config.LoadConfig("../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	InitDB(config.CONFIG)

	ret, _ := GetVoteCount(1, "2")

	log.Println(ret)
}
