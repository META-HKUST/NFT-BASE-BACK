package model

import (
	"NFT-BASE-BACK/base"
	"log"
)

var(
	queryInfoByTokenID = string("select item_name,item_data,label from items,item_label where items.item_id = item_label.item_id and items.item_id=?;")
)

type MetaData struct {
	Name		string	`json:"name" db:"item_name"`
	DataType	string	`json:"data_type" db:"label""`
	IpfsUrl		string	`json:"ipfs_url" db:"item_data"`
}
func GetMetaDataInfo(tokenId string) (MetaData,error){
	var MetaDataInfo MetaData
	err := db.Get(&MetaDataInfo,queryInfoByTokenID,tokenId)
	if err != nil {
		log.Println(base.QueryError, base.QueryError.String(), err)
		return MetaData{},err
	}
	return MetaDataInfo,nil
}
