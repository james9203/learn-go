package main

import (
	"time"
	"fmt"
)

func main() {
	//China doesn't have daylight saving.It uses a fixed 8 hour offset from UTC.
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)
}
