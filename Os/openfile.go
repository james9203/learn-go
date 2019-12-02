package main

import (
	"os"
	"log"
	"fmt"
)

func  openfilew()  {
	f,err := os.OpenFile("D:/reastdata/a.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil{
		log.Fatal(err)
	}
	if _,err := f.Write([]byte("appended some data\r\n")); err != nil{
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close();err != nil{
		log.Fatal(err)
	}

}

func  main()  {
	fmt.Println("-----------------------------------")
	openfilew()
	fmt.Println("-----------------------------------")
	f,err := os.OpenFile("D:/reastdata/a.txt",os.O_RDWR|os.O_CREATE,0755)

	if err != nil{
		log.Fatal(err)
	}

	if err := f.Close(); err != nil{
		log.Fatal(err)
	}
	fmt.Println("-----------------------------------")
}
