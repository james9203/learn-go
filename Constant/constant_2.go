package main

import "fmt"

const (
	A = "12"
	B = len(A)
	C = iota
)

func main()  {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
