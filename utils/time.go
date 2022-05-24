package utils

import (
	"log"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

func GetTimeNow() string {
	return time.Unix(time.Now().Unix(), 0).Format(timeLayout)
}

// compare if the given time is 15 minutes faster than current time
// @return true means more than 15minutes, false in contrast
func IsTimeValid(t string) bool {
	timeNow := time.Now()
	loc, _ := time.LoadLocation("Local")
	GivenTime, _ := time.ParseInLocation(timeLayout, t, loc)
	timeLoss := timeNow.Sub(GivenTime)
	log.Println("activate email token: time loss = ", timeLoss.Seconds())
	if timeLoss.Seconds() > 900 {
		return false
	} else {
		return true
	}
}
