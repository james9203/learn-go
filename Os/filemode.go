package main

import (
	"os"
	"log"
	"fmt"
)

func  main()  {
	fi,err := os.Lstat("D:/reastdata/a.txt")
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("permissions:%#o\n",fi.Mode().Perm())
	switch mode := fi.Mode();{
			case mode.IsRegular():
				fmt.Println("regular file")
			case mode.IsDir():
				fmt.Println("directory")
			case mode & os.ModeSymlink != 0:
				fmt.Println("symbolic link")
			case mode & os.ModeNamedPipe != 0:
				 fmt.Println("named pipe")
	}
}
