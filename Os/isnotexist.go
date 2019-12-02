package main

import (
	"os"
	"fmt"
)

func  main()  {

	filename := "D:/reastdata/b.txt"
	if _ ,err := os.Stat(filename);os.IsNotExist(err){
		fmt.Println("file does not exist!")
	}
}