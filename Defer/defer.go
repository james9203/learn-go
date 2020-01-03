package main

import (
	"fmt"
)

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	//fmt.Println(i)
	return i * 2
}
func increaseB() (r int) {
	defer func() {
		r++
	}()
	//fmt.Println(r)
	return 2 * r
}

//defer 是闭包引用，返回值被修改，所以 f1() 返回 1。
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

// 2.闭包引用，但是没有修改返回值 r
func f2() (r int) {
	t := 5
	// 1.赋值
	//r = t
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {

	// 1.赋值
	//r = 1
	// 2.r 作为函数参数，不会修改要返回的那个 r 值
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(increaseA())
		fmt.Println(increaseB())
	}
	fmt.Println("---------------------------")
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}
