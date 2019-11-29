package  main

import "fmt"

func relational()  {
	var a int = 21
	var b int = 10
	if(a==b){
		fmt.Printf("第一行-a 等于 b \n")
	}else {
		fmt.Printf("第一行 -a 不等于 b\n")
	}
	if(a<b){
		fmt.Printf("第二行- a 小于 b\n")
	}else {
		fmt.Printf("第二行 -a 不小于 b\n")
	}

	if(a>b){
		fmt.Printf("第三行- a 大于 b\n")
	}else{
		fmt.Printf("第三行- a 不大于 b\n")
	}

	a=5
	b=20
	if(a<=b){
		fmt.Printf("第四行 a 小于等于 b\n")
	}
	if(b>=a){
		fmt.Printf("第五行 b 大于等于 b\n")
	}

}

func logical()  {
	var a bool = true
	var b bool = false
	if(a&&b){
		fmt.Printf("第一行 - 条件为true \n")
	}
	if(a||b){
		fmt.Printf("第二行 - 条件为 true\n")
	}
	a=false
	b = true
	if(a&&b){
		fmt.Printf("第三行 - 条件为 true\n")
	}else{
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if(!(a&&b)){
		fmt.Printf("第四行 -条件为 true\n")
	}
}

func  bitwise()  {
	var a uint = 60 /*60=0011 1100*/
	var b uint = 13 /*13 = 0000 1101*/
	var c uint = 0
	c = a&b
	fmt.Printf("第一行 -c 的值为 %d\n",c)
	c = a|b
	fmt.Printf("第二行 -c 的值为 %d\n",c)
	c = a^b
	fmt.Printf("第三行 -c 的值为 %d\n",c)
	c = a<<2
	fmt.Printf("第四行 -c 的值为 %d\n",c)
	c = a>> 2
	fmt.Printf("第五行 -c 的值为 %d\n",c)
}

func  assignment()  {
	var a int = 21
	var c int

	c = a
	fmt.Printf("第一行 - = 运算符实例，c的值为=%d\n",c)
	c += a
	fmt.Printf("第二行 - += 运算符实例，c 的值为=%d\n",c)
	c -= a
	fmt.Printf("第三行好 - -= 运算符实例， c 的值为=%d\n",c)
	c *= a
	fmt.Printf("第四行好 - *= 运算符实例， c 的值为=%d\n",c)
	c /= a
	fmt.Printf("第五行好 - /= 运算符实例， c 的值为=%d\n",c)

	c = 200
	c <<=2
	fmt.Printf("第 六行  - <<= 运算符实例，c 值为 = %d\n", c )
	c >>= 2
	fmt.Printf("第 七行  - >>= 运算符实例，c 值为 = %d\n", c )
	c &= 2
	fmt.Printf("第 八行  - &= 运算符实例，c 值为 = %d\n", c )
	c ^= 2
	fmt.Printf("第 九行  - ^= 运算符实例，c 值为 = %d\n", c )
	c |= 2
	fmt.Printf("第 十行  - |= 运算符实例，c 值为 = %d\n", c )
}
func orther(){

	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

	/*  & 和 * 运算符实例 */
	ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a);
	fmt.Printf("*ptr 为 %d\n", *ptr);

}

func  main()  {
	fmt.Println("--------------------------------")
	orther()
	fmt.Println("--------------------------------")
	assignment()
	fmt.Println("--------------------------------")
	bitwise()
	fmt.Println("--------------------------------")
	logical()
	fmt.Println("--------------------------------")
	relational()
	fmt.Println("--------------------------------")
	var a int = 21
	var b int = 10
	var c int
	c = a+b
	fmt.Printf("第一行 -c 的值为 %d\n",c)
	c = a-b
	fmt.Printf("第二行 -c 的值为 %d\n",c)
	c= a*b
	fmt.Printf("第三行 -c 的值为 %d\n",c)
	c = a/b
	fmt.Printf("第四行 -c 的值为 %d\n",c)
	c = a%b
	fmt.Printf("第五行 -c 的值为 %d\n",c)
	a++
	fmt.Printf("第六行 -a 的值为 %d\n",a)
	a--
	fmt.Printf("第六行 -a 的值为 %d\n",a)
	fmt.Println("--------------------------------")

}