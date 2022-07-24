package model

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/entity"
	"log"
)

var (
	getMaxCollectionId = string("select max(collection_id) from collection")
	insertCollection   = string("insert into collection(collection_name ,logo_image ,feature_image ,banner_image ,items_count ,description ,owner, owner_name, created_at) values(?,?,?,?,?,?,?,?,?);")
	queryCollection    = string("select collection_id, collection_name ,logo_image ,feature_image ,banner_image ,items_count ,description ,owner, owner_name, created_at from collection where collection_id=?;")

	// edit collection
	UpdateCollectionName = string("update collection set collection_name=? where collection_id=?;")
	UpdateLogoImage      = string("update collection set logo_image=? where collection_id=?;")
	UpdateFeatureImage   = string("update collection set feature_image=? where collection_id=?;")
	UpdateBannerImage    = string("update collection set banner_image=? where collection_id=?;")
	UpdateDescription    = string("update collection set description=? where collection_id=?;")

	getCollectionName = string("select collection_name from collection where collection_id = ?")

	insertCollectionLable = string("insert into collection_label(collection_id, label) values(?,?);")
	queryCollectionLable  = string("select label from collection_label where collection_id=?")

	deleteCollection = string("DELETE FROM collection WHERE owner=?;")

	updateCoLabel = string("update collection_label set label=? where collection_id=?;")
)

type CollectionLabel struct {
	CollectionID    int    `db:"collection_id"`
	CollectionLabel string `db:"label"`
}

func DeleteCollection(UserId string) error {
	_, e1 := db.Exec(deleteCollection, UserId)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	return nil
}

func SearchCoLable(CoId interface{}) ([]string, error) {
	var LableSlice []string
	err := db.Select(&LableSlice, queryCollectionLable, CoId)
	if err != nil {
		return []string{}, err
	}
	return LableSlice, nil
}

func CreateCollectionLabel(label CollectionLabel) ([]string, error) {
	_, err := db.Exec(insertCollectionLable,
		label.CollectionID,
		label.CollectionLabel,
	)
	if err != nil {
		log.Println(err)
		return []string{}, err
	}
	var ret []string
	err = db.Select(&ret, queryCollectionLable, label.CollectionID)
	if err != nil {
		return []string{}, err
	}
	log.Println("collection label: ", ret)
	return ret, nil
}

func CreatCollection(collection_name string, logo_image string, feature_image string, banner_image string, items_count int, description string, owner string, owner_name string, created_at string) error {
	owner_name, _ = GetUserName(owner)
	log.Println("create collection owner_name: ", owner_name)
	if owner_name == "" {
		owner_name = owner
	}
	_, e := db.Exec(insertCollection, collection_name, logo_image, feature_image, banner_image, items_count, description, owner, owner_name, created_at)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return e
	}
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

func EditCollectionLable(label []string, CollectionId int) error {

	labelss, _ := SearchCoLable(CollectionId)
	for i := 0; i < len(label); i++ {
		contain := false
		for j := 0; j < len(labelss); j++ {
			if labelss[j] == label[i] {
				contain = true
			}
		}
		if contain == false {
			_, err := db.Exec(insertCollectionLable,
				CollectionId,
				label[i],
			)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}

type CoAndLabel struct {
	entity.Collection
	Label []string `json:"label" db:"label"`
}

func GetCoAndLabel(collectionId interface{}) (CoAndLabel, error) {
	Co := CoAndLabel{}
	var err error
	Co.Collection, err = GetCollection(collectionId)
	if err != nil {
		log.Println(err)
		return CoAndLabel{}, err
	}
	Co.Label, err = SearchCoLable(collectionId)

	log.Println("collection_id: ", collectionId, "labels: ", Co.Label)
	if err != nil {
		log.Println(err)
		return CoAndLabel{}, err
	}
	return Co, nil
}

func GetCollection(collectionId interface{}) (entity.Collection, error) {
	var g entity.Collection
	err := db.Get(&g, queryCollection, collectionId)
	if err != nil {
		log.Println(err)
		return entity.Collection{}, err
	}
	return g, nil
}

func GetCollectionName(collectionId interface{}) (string, error) {
	var g string
	err := db.Get(&g, getCollectionName, collectionId)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return g, nil
}
