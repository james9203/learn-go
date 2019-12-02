package main

import (
	"fmt"
	"os"
)

func  mapper( placeholderName string)  string {
	switch placeholderName {
	case "DAY_PART":
			return  "morning"
	case "NAME":
			return  "Gopher"
	}
	return  ""
}

func  main()  {
	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper))
}