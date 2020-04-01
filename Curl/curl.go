package main

import (
	"github.com/mikemintang/go-curl"
	"log"
	"fmt"
)

func main() {

	url := ""

	cli := goz.NewClient()
	header := make(map[string]interface{})
	header["Content-Type"] = "application/json"
	resp, err := cli.Post(url, goz.Options{
		Headers:     header,
		HeaderLower: false,
		JSON: map[string]interface{}{
			"headId":   "a06751fd-1ff9-4799-be51-011f95cfccf8",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	res, _ := resp.GetBody()
	fmt.Println(res)
	headers := resp.GetRequest().Header
	fmt.Println(headers)
	// Output: [Bar Baz]

}
