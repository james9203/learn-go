package main

import (
	"fmt"
	"unsafe"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func  iotas()  {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f=100
		g
		h= iota
		i
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}

func  iotas2()  {
	const (
		i=1<<iota
		j=3<<iota
		k
		l
	)
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}

func  main()  {
	iotas()
	iotas2()
	fmt.Println(a,b,c)
	const  LENGHT int  = 10
	const  WIDTH int   =  5

	var area int

	const  a,b,c  = 1,false,"str" //多重复值

	area = LENGHT * WIDTH

	fmt.Printf("面积为：%d",area)
	fmt.Println()
	fmt.Println(a,b,c)
}
