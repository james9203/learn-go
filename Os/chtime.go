package main

import (
	"time"
	"os"
	"log"
	"fmt"
)

func main()  {
	mtime := time.Date(2006,time.February,1,3,4,5,0,time.UTC)
	atime := time.Date(2007,time.March,2,4,5,6,0,time.UTC)
	if err := os.Chtimes("D:/reastdata/a.txt",atime,mtime);err != nil{
		log.Fatal(err)
	}else {
		fmt.Println("success!")
	}
}
