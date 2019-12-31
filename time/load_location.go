package main

import (
	"time"
	"fmt"
)

func main() {
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	timeInUTC := time.Date(2019, 12, 31, 18, 6, 0, 0, time.UTC)
	fmt.Println(timeInUTC.In(location))
}
