package main

import "fmt"

func  arraypointer()  {
	const  MAX int  = 3
	a := []int{10,100,200}
	var i int
	var ptr [MAX]*int
	for i=0;i<MAX ;i++  {
		ptr[i] = & a[i]
	}
	for  i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
	}
}

func  main()  {
	fmt.Println("-------------------------")
	var  a int  =  10
	fmt.Printf("变量的地址为：%x\n",&a)
	fmt.Println("-------------------------")
	var ip *int
	ip = &a
	fmt.Printf("变量的地址为：%x\n",&a)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
	fmt.Println("-------------------------")

	var  ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr  )
	fmt.Println("-------------------------")
	if(ptr==nil){
		fmt.Println("ptr 的指针为空")
	}
	fmt.Println("-------------------------")
	arraypointer()
	fmt.Println("-------------------------")
	
}
