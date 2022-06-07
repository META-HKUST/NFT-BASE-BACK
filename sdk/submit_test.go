package sdk

import (
	"testing"
)

func TestSubmit(t *testing.T) {
	submitTests := []struct {
		username     string
		contractName string
		args         []string
	}{
		{
			username:     "zwang",
			contractName: "PublicMint",
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
