package main

import (
	"fmt"
	"time"
)

func main()  {
	ChanGo()
	fmt.Println("----------------------------------")
	go Go()
	time.Sleep(2*time.Second)
}

func ChanGo()  {
	c := make(chan  bool)
		go func() {
			fmt.Println("go go go go go go go go ")
			c <- true
		}()
	<-c
}

func Go()  {
	fmt.Println("GO GO GO GO Go")
}
