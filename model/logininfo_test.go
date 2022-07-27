package model

import (
	"NFT-BASE-BACK/config"
	"log"
	"testing"
)

func TestUpdateVerifyCode(t *testing.T) {
	err := config.LoadConfig("../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	InitDB(config.CONFIG)
	err = UpdateVerifyCode("mingzheliu@ust.hk", "123456", "2022/07/27 22:01:18")
	if err != nil {
		log.Fatal("cannot update verify code", err)
	}

	log.Println("Update Verify Code Success")
}
