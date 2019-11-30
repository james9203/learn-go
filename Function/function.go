package main

import "fmt"

func  max(num1,num2 int) int {
	if (num1>num2){
		return  num1
	}else{
		return  num2
	}
}

func  swap(x,y string) (string,string)  {
	return  y,x
}

func  main()  {
	fmt.Println("----------------------------------------")
	var a int = 100
	var  b int   =  200
	var ret int
	ret = max(a,b)
	fmt.Printf("最大值是 %d\n",ret)
	fmt.Println("----------------------------------------")

	c,d := swap("google" ,"baidu")
	fmt.Println(c,d)
	fmt.Println("----------------------------------------")


}