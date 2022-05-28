package sdk

import (
    "testing"
)

func TestEnroll(t *testing.T) {
    var enrollTests = []struct {
        username string
    }{
        {"zwang"},
        {"zzding4"},
        {"zzding5"},
        {"zzding6"},
    }

    for _, v := range enrollTests {
        result, err := Enroll(v.username)
        if err != nil {
            t.Log(err)
        }
        t.Log(result)

    }
}
