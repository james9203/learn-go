package main

import (
	"fmt"
	"strconv"
	"log"
)

type (
	byte1  int8
	rune   int32
	ByteSize int64
	bo   bool
)
func main()  {
	var a  byte1
	var b rune
	var c ByteSize
	var d bo
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//var e,f,g,h int = 1,2,3,4
	e,f,g,h := 1,2,3,4
	fmt.Println(e,f,g,h)
	var j float32 = 1.1
	k := int(j)
	fmt.Println(k)
	var l int = 65
	n := strconv.Itoa(l)
	m := string(l)
	fmt.Println(m)
	str,err :=strconv.Atoi(n)
	if(err==nil){
		fmt.Println(str)
	}else{
		log.Fatal(err)
	}

}
