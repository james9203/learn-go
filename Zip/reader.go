package main

import (
	"archive/zip"
	"log"
	"fmt"
	"io"
	"os"
)

func  main()  {
	//Open a zip archive for reading .
	r,err := zip.OpenReader("D:/reastdata/readme.zip")
	if err != nil{
		log.Fatal(err)
	}
	defer  r.Close()
	for _,f := range r.File {
			fmt.Printf("Contens of %s:\n",f.Name)
			rc,err := f.Open()
			if err != nil{
				log.Fatal(err)
			}
			_,err = io.CopyN(os.Stdout,rc,68)
			if err != nil {
				log.Fatal(err)
			}
			rc.Close()
			fmt.Println()
	}
}
