package utils

import (
	"log"
	"testing"
)

func TestGenEmail(t *testing.T) {
	//err := Email("Mingzhe", "mingzheliu@ust.hk", "123")
	//if err != nil {
	//	log.Println(err)
	//}

	err := Email("Mingzhe", "1721062927@qq.com", "123")
	if err != nil {
		log.Println(err)
	}
}
