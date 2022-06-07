package sdk

import (
    "testing"
)

func TestEnroll(t *testing.T) {
    var enrollTests = []struct {
        username string
    }{
        {"zwang"},
    }

    for _, v := range enrollTests {
        result, err := Enroll(v.username)
        if err != nil {
            t.Log(err)
        }
        t.Log(result)

    }
}
