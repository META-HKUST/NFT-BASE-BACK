package v2

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/config"
	"NFT-BASE-BACK/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
)

func TestEditItem(t *testing.T) {
	err := config.LoadConfig("../../../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	db, err := sqlx.Open(config.CONFIG.DBDriver, config.CONFIG.DBSource)
	if err != nil {
		log.Println(base.OpenSqlError, base.OpenSqlError.String(), err)
		return
	}

	err2 := db.Ping()
	if err2 != nil {
		log.Println(base.ConnectSqlError, base.ConnectSqlError.String(), err2)
		return
	}
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(20)


	code := model.EditItem("1010","hammer","update","",[]string{"123","allen"})
	if code != nil {
		fmt.Println(code)
	}

	itemInfo,code := model.GetItemInfo("1010")
	if code != nil{
		fmt.Println(code)
	}

	fmt.Println(itemInfo)
}


func TestGetItemList(t *testing.T) {
	err := config.LoadConfig("../../../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	db, err := sqlx.Open(config.CONFIG.DBDriver, config.CONFIG.DBSource)
	if err != nil {
		log.Println(base.OpenSqlError, base.OpenSqlError.String(), err)
		return
	}

	err2 := db.Ping()
	if err2 != nil {
		log.Println(base.ConnectSqlError, base.ConnectSqlError.String(), err2)
		return
	}
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(20)


	iteminfo,code := model.GetItemList(1,2,"ffff",false,false,false,"","nft",false,false,0)
	if code != nil {
		fmt.Println(code)
	}
	fmt.Println(len(iteminfo))
	fmt.Println(iteminfo)
}

//func TestGetProfiles(t *testing.T) {
//	err := config.LoadConfig("../../../config")
//	if err != nil {
//		log.Fatal("cannot load config", err)
//	}
//
//	db, err := sqlx.Open(config.CONFIG.DBDriver, config.CONFIG.DBSource)
//	if err != nil {
//		log.Println(base.OpenSqlError, base.OpenSqlError.String(), err)
//		return
//	}
//
//	err2 := db.Ping()
//	if err2 != nil {
//		log.Println(base.ConnectSqlError, base.ConnectSqlError.String(), err2)
//		return
//	}
//	db.SetMaxOpenConns(100)
//	db.SetConnMaxIdleTime(20)
//	model.GetUserListByKey(db,"en")
//}