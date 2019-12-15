package main

import "fmt"

func printSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
func  slices()  {
	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}
func  slices_2()  {

	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1,numbers)
	printSlice(numbers1)
}

func slices_3()  {
	var s1 []int
	a := [10]int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(a)
	s2 := a[5:]
	fmt.Println("-------------------------------")
	fmt.Println(s2)
	fmt.Println("-------------------------------")
	fmt.Println(s1)
	fmt.Println("-------------------------------")
	s3 := make([]int,3,10)
	fmt.Println(s3)

	c := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sc := c[2:5]
	fmt.Println(len(sc),cap(sc))
	sb := sc[1:3]
	sd  := append(sc,1,2,3,4,5,5)

	fmt.Println(string(sb))
	fmt.Println(sd)
	fmt.Println("-------------------------------")
	s4 := []int{1,2,3,4,5,6}
	s5 := []int{7,8,9}
	copy(s4,s5)
	fmt.Println(s4)

}
func  main()  {
	fmt.Println("-------------------------------")
	slices_3()
	fmt.Println("-------------------------------")
	var number = make([]int,3,5)
	printSlice(number)
	fmt.Println("-------------------------------")
	var numbers []int

	printSlice(numbers)

	if(numbers == nil){
		fmt.Printf("切片是空的\n")
	}
	fmt.Println("-------------------------------")
	slices()
	fmt.Println("-------------------------------")
	slices_2()
	fmt.Println("-------------------------------")

}
