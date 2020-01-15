package main

import (
	"flag"
	"io/ioutil"
	"fmt"
)

func main() {
	fptr := flag.String("fpath", "text.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
