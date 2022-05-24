package model

import (
	"NFT-BASE-BACK/base"
	"fmt"
	"log"
)
type TokenInfo struct {
	TokenID		int64
	Url			string
}
var(
	//NFT data storage related sentences
	addTokenUrl = string("insert into nftstore(tokenid,url) values(?,?);")
	findUrlByID = string("select * from nftstore where tokenid=?;")
)

// Store URL and tokenid to database
func StoreUrl(tokenId int64,url string) base.ErrCode {
	result, err := db.Exec(addTokenUrl,tokenId,url)
	if err != nil {
		log.Println(base.InsertError, base.InsertError.String(), err)
		return base.InsertError
	}
	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	log.Println("rowsAffected: ", rowsAffected, "lastInsertId: ", lastInsertId)
	return base.Success
}

func GetUrlByTokenId(tokenId int64) (TokenInfo,error){
	tokenInfo := TokenInfo{}
	fmt.Println(tokenId)
	err := db.Get(&tokenInfo,findUrlByID,tokenId)
	if err != nil {
		log.Println(err)
		return TokenInfo{},err
	}
	return tokenInfo,nil
}
