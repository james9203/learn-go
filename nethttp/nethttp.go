package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func hGet()  {
	resp,err := http.DefaultClient.Get("http://www.baidu.com")
	defer  resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data ,err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(data))
	}
}

func main()  {
	hGet()

}