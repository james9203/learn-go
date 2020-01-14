package main

import "fmt"

type add func(a int, b int) int

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func an_a() {
	a := func() {
		fmt.Println("hello world first class function!")
	}
	a()
	fmt.Printf("%T\n", a)
}
func an_b() {
	func() {
		fmt.Println("hello world first class function!")
	}()
	a := "Gophers"
	func(n string) {
		fmt.Println("Welcome", n)
	}(a)
}

func an_c() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(6, 5)
	fmt.Println("Sum", s)
}

func simple(a func(a int, b int) int) {
	fmt.Println(a(60, 7))
}

func simpleb() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}
func an_d() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}
func an_student() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	c := filter(s, func(s student) bool {
		if s.country == "India" {
			return true
		}
		return false
	})
	fmt.Println(f)
	fmt.Println(c)
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func an_e() {
	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 9
	})
	fmt.Println(r)
}

func main() {
	an_a()
	an_b()
	an_c()
	f := func(a, b int) int {
		return a + b
	}
	s := simpleb()
	fmt.Println(s(60, 7))
	simple(f)
	an_d()
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
	an_student()
	an_e()

}
