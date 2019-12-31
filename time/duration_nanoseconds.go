package main

import (
	"time"
	"fmt"
)

func main() {
	u, _ := time.ParseDuration("1µs")
	fmt.Printf("One microsecond is %d nanoseconds.\n", u.Nanoseconds())
}
