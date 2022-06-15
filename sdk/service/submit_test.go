package service

import (
	"NFT-BASE-BACK/config"
	"log"
	"testing"
)

func TestSubmit(t *testing.T) {

	err := config.LoadConfig("../../config")
	if err != nil {
		log.Fatal(err)
	}

	submitTests := []struct {
		username     string
		contractName string
		args         []string
	}{
		{
			username:     "zwang",
			contractName: "PublicMint",
			args:         []string{},
		},
	}

	for _, v := range submitTests {
		result, err := Submit(v.username, v.contractName, v.args...)
		if err != nil {
			t.Log(err)
		}
		t.Log(result)
	}

}
