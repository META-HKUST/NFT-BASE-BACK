package model

import (
	"NFT-BASE-BACK/base"
	"log"
)

var (
	insertCollection = string("insert into collection(collection_name ,logo_image ,feature_image ,banner_image ,items_count ,description ,owner, creater, created_at) values(?,?,?,?,?,?,?,?,?);")
)

func CreatCollection(collection_name string, logo_image string, feature_image string, banner_image string, items_count int, description string, owner string, creater string, created_at string) error {
	result, e := db.Exec(insertCollection, collection_name, logo_image, feature_image, banner_image, items_count, description, owner, creater, created_at)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return e
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	log.Println("rowsAffected: ", rowsAffected, "lastInsertId: ", lastInsertId)
	return nil
}
