package main

import (
	"encoding/json"
	"fmt"
)

type Domainrecord struct {
	Id          int    `json:"id"`
	Domain_id   int    `json:"domain_id"`
	Name        string `json:"name"`
	Record_type string `json:"record_type"`
	Content     string `json:"content"`
	Ttl         int    `json:"ttl"`
	Prio        int    `json:"prio"`
}

func main() {

	record := Domainrecord{}
	record.Id = 10
	record.Domain_id = 100
	record.Name = "lotus668.com"
	record.Record_type = "SOA"
	record.Content = "10.10.10.10"
	record.Ttl = 30
	record.Prio = 20
	rv, errorr := json.Marshal(record)
	fmt.Println(string(rv))
	fmt.Println(errorr)
	fmt.Println("------------------")
	return

	type UserInfo struct {
		Id      int
		Name    string
		Address string
	}

	type Comiket struct {
		id    int
		title string
	}

	type Test struct {
		Id      int    `json:"id"`
		Content string `json:"content"`
		IsDel   bool   `json:",string"`
		Type    int    `json:"type, omitempty"`
	}

	c := Comiket{1, "AS动漫游戏嘉年华"}
	v, error := json.Marshal(c)
	fmt.Println(string(v))
	fmt.Println(error)
	fmt.Println("------------------")

	User := UserInfo{1, "taylor", "USA"}
	b, err := json.Marshal(User)
	fmt.Println(string(b))
	fmt.Println(err)
	fmt.Println("------------------")

	t := Test{Id: 1, Content: "hello world", IsDel: true}
	val, _ := json.Marshal(t)
	fmt.Println(string(val))
}
