package main

import (
	"unsafe"
	"fmt"
)

func main() {

	var p float64 = 99
	var c int
	var e string
	var g map[int] string
	a := unsafe.Sizeof(p)
	d := unsafe.Sizeof(c)
	f := unsafe.Sizeof(e)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(a)

}
