package main

import (
	"os"
	"fmt"
)

func  main()  {
	os.Setenv("Name","gopher")
	os.Setenv("BURROW","/usr/gopher")
	fmt.Printf("%s lives in %s.\n",os.Getenv("NAME"),os.Getenv("BURROW"))
}
