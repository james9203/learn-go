package main

import (
	"fmt"
)

func getName(params interface{})  {
	var stringSlice []interface{}
	stringSlice = append(stringSlice, params)
	stringSlice = append(stringSlice, params)

	fmt.Println(stringSlice)
}

func main() {
	getName("redis")

}
