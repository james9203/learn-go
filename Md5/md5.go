package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"crypto/md5"
)

func md5New() {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")  // 加密字符串
	io.WriteString(h, "And Leon's getting laaarger!") //加密字符串
	fmt.Printf("%x\n", h.Sum(nil))
}

//read file content and md5
func md5File() {
	f, err := os.Open("d:/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}

//use md5.sum create md5
func md5Byte() {
	data := []byte("The fog is getting thicker!And Leon's getting laaarger!")
	fmt.Printf("%x\n", md5.Sum(data))
}

func main() {
	fmt.Println("-------------------------")
	md5New()
	fmt.Println("-------------------------")
	md5File()
	fmt.Println("-------------------------")
	md5Byte()
	fmt.Println("-------------------------")

}
