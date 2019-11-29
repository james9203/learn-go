package main

import "fmt"

func main() {
	//声明并初始化
	var a string = "Hello,World!"
	fmt.Println(a)
	var b, c = 1, 2
	fmt.Println(b, c)
	//声明没有初始化
	var d int
	fmt.Println(d)
	var e bool
	fmt.Println(e)

	//nil
	var f *int //指针
	fmt.Println(f)
	var g []int //数组
	if (g == nil) {
		fmt.Println(g)
	}
	var h map[string]int
	if (h == nil) {
		fmt.Println(h)
	}
	var i chan int
	if (i == nil) {
		fmt.Println(i)
	}
	var j func(string) int
	if (j == nil) {
		fmt.Println(j)
	}
	var k error // error 是接口

	if (k == nil) {
		fmt.Println(k)
	}

	var s string
	var l int
	var m float64
	fmt.Printf("%v %v %q\n",l,m,s)

	intval := 1

	intval, stringval := 1,"222"
	fmt.Println(intval)
	fmt.Println(stringval)

}
