package main

import (
	"time"
	"fmt"
)

func main() {
	m, _ := time.ParseDuration("1m30s")
	fmt.Printf("Take off t-%.0f seconds.", m.Seconds())
}
