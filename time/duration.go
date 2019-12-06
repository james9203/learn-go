package main

import (
	"time"
	"fmt"
)

func expensiveCall()  {
	for a:=1;a<10000500 ;a++  {

	}
}

func main()  {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	//fmt.Println(t0)
}
