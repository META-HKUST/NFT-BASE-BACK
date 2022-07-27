package utils

import (
	"NFT-BASE-BACK/config"
	"log"
	"testing"
)

func TestTimeNow(t *testing.T) {
	err := config.LoadConfig("../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	time := GetTimeNow()
	log.Println(time)
	log.Println(time)
	log.Println(time)
	log.Println(time)
	log.Println(time)
}
