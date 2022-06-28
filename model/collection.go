package model

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/entity"
	"log"
)

var (
	getMaxCollectionId = string("select max(collection_id) from collection")
	insertCollection   = string("insert into collection(collection_name ,logo_image ,feature_image ,banner_image ,items_count ,description ,owner, creater, created_at) values(?,?,?,?,?,?,?,?,?);")
	queryCollection    = string("select collection_id, collection_name ,logo_image ,feature_image ,banner_image ,items_count ,description ,owner, creater, created_at from collection where collection_id=?;")

	// edit collection
	UpdateCollectionName = string("update collection set collection_name=? where collection_id=?;")
	UpdateLogoImage      = string("update collection set logo_image=? where collection_id=?;")
	UpdateFeatureImage   = string("update collection set feature_image=? where collection_id=?;")
	UpdateBannerImage    = string("update collection set banner_image=? where collection_id=?;")
	UpdateDescription    = string("update collection set description=? where collection_id=?;")
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

func GetMaxCollectionId() (int, error) {
	var g int
	err := db.Get(&g, getMaxCollectionId)
	if err != nil {
		return -1, err
	}
	return g, nil
}

func EditCollectionName(Arg string, CollectionId int) error {
	_, e := db.Exec(UpdateCollectionName, Arg, CollectionId)
	if e != nil {
		return e
	}
	return nil
}

func EditLogoImage(Arg string, CollectionId int) error {
	_, e := db.Exec(UpdateLogoImage, Arg, CollectionId)
	if e != nil {
		return e
	}
	return nil
}

func EditFeatureImage(Arg string, CollectionId int) error {
	_, e := db.Exec(UpdateFeatureImage, Arg, CollectionId)
	if e != nil {
		return e
	}
	return nil
}

func EditBannerImage(Arg string, CollectionId int) error {
	_, e := db.Exec(UpdateBannerImage, Arg, CollectionId)
	if e != nil {
		return e
	}
	return nil
}

func EditDescription(Arg string, CollectionId int) error {
	_, e := db.Exec(UpdateDescription, Arg, CollectionId)
	if e != nil {
		return e
	}
	return nil
}

func GetCollection(collectionId int) (entity.Collection, error) {
	var g entity.Collection
	err := db.Get(&g, queryCollection, collectionId)
	if err != nil {
		log.Println(err)
		return entity.Collection{}, err
	}
	return g, nil
}
