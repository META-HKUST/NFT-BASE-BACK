package service

import (
	"NFT-BASE-BACK/fileservice"
	"NFT-BASE-BACK/utils"
	"io"
	"log"
)

func Upload(key string,content io.Reader) (string,string,error){
	_, url, err := fileservice.Upload(key,content)
	if err != nil {
		log.Println(err)
		return "","",err
	}
	encryText,err := utils.AesEcpt.AesBase64Encrypt(url.String())
	//encryText,err:= utils.Encrypt(url.String(),fileservice.COSCONFIG.CryptoKey)
	if err != nil {
		return "","",err
	}
	return url.String(),encryText,nil
}
