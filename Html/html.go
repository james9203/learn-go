package main

import (
	"fmt"
	"html"
)

func main()  {
	 s := `"Fran & Freddie's Diner" <tasty@example.com>`
	 es := html.EscapeString(s)
	 fmt.Println(es)
	 us := html.UnescapeString(es)
	fmt.Println(us)

}