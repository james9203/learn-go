package main

import "fmt"

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error:first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error:last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullname")
}
func main() {
	defer fmt.Println("deferred call in main")
	firstname := "Elon"
	fullName(&firstname, nil)
	fmt.Println("returned normally form main")
}
