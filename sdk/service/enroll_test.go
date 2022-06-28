package service

import (
	"NFT-BASE-BACK/config"
	"log"
	"testing"
)

func TestEnroll(t *testing.T) {
	err := config.LoadConfig("../../config")
	if err != nil {
		log.Fatal(err)
	}

	var enrollTests = []struct {
		username string
	}{
		{"zwang"},
	}

	for _, v := range enrollTests {
		err := Enroll(v.username)
		if err != nil {
			t.Log(err)
		}

	}
}
