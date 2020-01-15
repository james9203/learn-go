package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	data, err := ioutil.ReadFile("d:/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
