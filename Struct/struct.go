package main

import "fmt"

type person struct {
	Name string
	Age  int
}
type Human struct {
	Sex int
}

type  Teacher struct {
	Human
	Name string
	Age int
}

func main() {
	a := person{
		Name: "james",
		Age:  28,
	}
	b := person{
		Name: "james",
		Age:  39,
	}
	fmt.Println(a==b)
	//a.Name = "james"
	//a.Age = 28
	A(a)
	fmt.Println(a)
	B(&a)
	fmt.Println(a)
	c := Teacher{Name:"james",Age:20,Human:Human{Sex:1}}
	fmt.Println(c)
	c.Sex = 0
	fmt.Println("-------------------------------------")
	fmt.Println(c)
}

func A( per person)  {
	per.Age = 13
	fmt.Println("A",per)
}
func B(per *person)  {
	per.Age = 12
	fmt.Println("B",per)
}
