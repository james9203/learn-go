package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId    int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) string {
	var query string = ""
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query = fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s,%d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s,\"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("unsupported type")
			}
		}
		query = fmt.Sprintf("%s)", query)
		return query
	}
	return query
}

func createOrderQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d,%d)", o.orderId, o.customerId)
	return i
}

func reflect_a() {
	i := 10
	fmt.Printf("%d %T\n", i, i)
}
func reflect_b() {
	o := order{
		orderId:    10,
		customerId: 617,
	}
	fmt.Println(o)
	fmt.Println(createOrderQuery(o))
}
func reflect_c() {
	o := order{
		orderId:    10,
		customerId: 617,
	}
	fmt.Println(createQuery(o))
	e := employee{
		name:    "james",
		id:      565,
		address: "Coimbatore",
		salary:  80000,
		country: "India",
	}
	fmt.Println(createQuery(e))
}
func main() {
	fmt.Println("------------------------")
	reflect_a()
	fmt.Println("------------------------")
	reflect_b()
	fmt.Println("------------------------")
	reflect_c()
	fmt.Println("------------------------")

}
