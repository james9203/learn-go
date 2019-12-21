package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := 'a'
	invalid := rune(0xfffffff)
	fmt.Println(utf8.ValidRune(valid))
	fmt.Println(utf8.ValidRune(invalid))
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('界'))
	buf := []byte("a界")
	fmt.Println(utf8.RuneStart(buf[0]))
	fmt.Println(utf8.RuneStart(buf[1]))
	fmt.Println(utf8.RuneStart(buf[2]))
	bufa := []byte{228, 184, 150} // 世
	fmt.Println(utf8.FullRune(bufa))
	fmt.Println(utf8.FullRune(bufa[:2]))
	bufc := []byte("Hello, 世界")
	fmt.Println("bytes =", len(bufc))
	fmt.Println("runes =", utf8.RuneCount(bufc))
	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.RuneCountInString(str))
}
