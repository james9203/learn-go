package main

import (
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"text/scanner"
	"os"
)

func main() {
	const input = "1234 5678 12345678901234567890"

	sacnner := bufio.NewScanner(os.Stdin)

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	sacnner.Split(split)
	for sacnner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := sacnner.Err(); err != nil {
		fmt.Printf("Invalid input:%s", err)
	}

}
