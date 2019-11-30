package main

import (
	"fmt"
	"reflect"
)

func  ifsentence()  {
	var a int = 100
	var b int = 200

	if a==100{
		if b==200 {
			fmt.Printf("a的值为100，b 的值为200\n")
		}
	}
	fmt.Printf("a值为：%d\n",a)
	fmt.Printf("b 值为：%d\n",b)

	if a<20 {
		fmt.Printf("a 小于20 \n")
	}else{
		fmt.Printf("a 不小于20 \n")
	}
	fmt.Printf("a 的值为： %d\n",a)

}

func  switchsentence(){
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70 : grade= "C"
	default: grade = "D"
	}
	switch {
	case grade == "A" :
		fmt.Printf("优秀!\n" )
	case grade == "B", grade == "C" :
		fmt.Printf("良好\n" )
	case grade == "D" :
		fmt.Printf("及格\n" )
	case grade == "F":
		fmt.Printf("不及格\n" )
	default:
		fmt.Printf("差\n" );
	}
	fmt.Printf("你的等级是 %s\n", grade );

}

func  selectsentence(){
	var c1 ,c2,c3 chan int
	var i1 ,i2 int
	select {
		case i1 = <-c1:
			fmt.Printf("received ",i1,"from c1\n")
	     case  c2 <-i2:
	     	fmt.Printf("sent ",i2,"to c2\n")
		case i3 ,ok := (<-c3):
			if ok{
				fmt.Printf("received",i3,"from c3\n")
			}else {
				fmt.Printf("c3 is closed\n")
			}
		default:
			fmt.Printf("no communication \n")

	}
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func loopsentence() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("---------------------------")
	sum = 1
	for ; sum <= 10; {
		sum += sum;
	}
	fmt.Println(sum)
	fmt.Println("---------------------------")
	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
	fmt.Println("---------------------------")
	sum = 0
	for  {
		sum++
		if(sum==100){
			break
		}
	}
	fmt.Println(sum)
	fmt.Println("---------------------------")

	strings := []string{"google","baidu"}
	number := []int{1,2,3,4,5,6,8899,10}
	fmt.Println(typeof(strings))
	fmt.Println(typeof(number))
	for i,s := range  strings{
		fmt.Println(i,s)
	}
	fmt.Println("---------------------------")

	numbers := [6]int{1, 2, 3, 5}
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	fmt.Println("---------------------------")


}

func main() {
	fmt.Println("---------------------------------------")
	selectsentence()
	fmt.Println("---------------------------------------")
	switchsentence()
	fmt.Println("---------------------------------------")
	ifsentence()
	fmt.Println("---------------------------------------")
	loopsentence()
	fmt.Println("---------------------------------------")
}
