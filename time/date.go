package main

import (
	"time"
	"fmt"
)

func main() {
	a := time.Unix(1579256205,0)
	fmt.Println(a)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}
