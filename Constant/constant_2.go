package main

import "fmt"

const (
	AA = "12"
	BB = len(AA)
	CC = iota
)

const (
	B float64 = 1<<(iota*10)
	KB
	MB
	GB
)

func main()  {
	fmt.Println(AA)
	fmt.Println(BB)
	fmt.Println(CC)
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

	var a int = 1
	var p *int = &a
	fmt.Println(p)
	fmt.Println(*p)

}
