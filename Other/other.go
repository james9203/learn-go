package main

import "fmt"

func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	return source * 2
}

func main() {
	fmt.Println(doubleScore(0))
	fmt.Println(doubleScore(20.0))
	fmt.Println(doubleScore(50.0))
}
