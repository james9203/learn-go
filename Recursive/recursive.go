package main

import "fmt"

func Factorial( n uint64)(result uint64)  {
	if(n>0){
		result = n * Factorial(n-1)
		return result
	}
	return  1
}
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main()  {
	fmt.Println("---------------------------------------")
	var i uint64 = 10
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))
	fmt.Println("---------------------------------------")
	var a int
	for a=0; a<10; a++ {
		fmt.Printf("%d\t", fibonacci(a))
	}
	fmt.Println("---------------------------------------")
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)
}
