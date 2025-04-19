package helper

import (
	"time"
)

func GetTimeNow() time.Time {
	// Init the location
	loc, _ := time.LoadLocation("Asia/Jakarta")

	// Set timezone and get now time
	return time.Now().In(loc)
}
