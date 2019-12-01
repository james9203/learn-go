package main

import (
	"os"
	"log"
	"fmt"
)

func  main()  {
	if err := os.Chmod("D:/reastdata/a.txt",0644); err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("success!")
	}
}
