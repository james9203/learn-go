package  main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func sorta()  {
	a := [...]int{5,2,6,3,9}
	num := len(a)
	for i:=0;i<num;i++ {
		for j:= i+1;j<num;j++{
			if a[i] < a[j]{
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}

func main()  {
	sorta()
	fmt.Println("---------------------------------------")
	var c [2] int
	var b [2]  int
	d := [...] int {0:12,1:2}
	b = c
	fmt.Println(b)
	fmt.Println(d)
	var dp *[2]int  = &d
	cp  := new([10] int)
	fmt.Println(cp)
	fmt.Println(dp)
	var n [10]int
	var i,j int
	for i=0;i<10 ;i++  {
		n[i]  = i+100
	}
	for j=0;j<10;j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}
	fmt.Println("-------------------------------------")
	var s [10]string
	for i=0;i<10;i++ {
		s[i] = "hello world"+string(strconv.Itoa(rand.Int()))
		fmt.Println(s[i])
	}
	fmt.Println("-------------------------------------")
	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}

	/* 输出数组元素 */
	for  i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
		}
	}
	fmt.Println("-------------------------------------")
}