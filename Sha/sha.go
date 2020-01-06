package main

import (
	"crypto/sha1"
	"io"
	"fmt"
	"os"
	"log"
	"crypto/sha256"
	"crypto/sha512"
)

func sha_1() {
	h := sha1.New()
	io.WriteString(h, "His money is twice tainted")
	io.WriteString(h, "taint yours and 'taint mine.")
	fmt.Printf("%x\n", h.Sum(nil))
}

func sha_1file() {
	f, err := os.Open("d:/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}

func sha_256() {
	h := sha256.New()
	io.WriteString(h, "hello world\n")
	//io.WriteString(h, "His money is twice tainted")
	//io.WriteString(h, "taint yours and 'taint mine.")
	fmt.Printf("%x\n", h.Sum(nil))
}
func sha_256file() {
	f, err := os.Open("d:/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}
func sha_512() {
	h := sha512.New()
	io.WriteString(h, "hello world\n")
	fmt.Printf("%x\n", h.Sum(nil))
}

func sha_512file() {
	f, err := os.Open("d:/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h := sha512.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}
func main() {
	fmt.Println("-----------------------------")
	sha_1()
	fmt.Println("-----------------------------")
	sha_1file()
	fmt.Println("-----------------------------")
	sha_256()
	fmt.Println("-----------------------------")
	sha_256file()
	fmt.Println("-----------------------------")
	sha_512()
	fmt.Println("-----------------------------")
}
