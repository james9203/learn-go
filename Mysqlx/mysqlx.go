package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	Db, err := sqlx.Open("mysql", "root:jamesroot@tcp(10.20.60.188:3306)/smart_wiki")
	if err != nil {
		fmt.Println("connect to mysql failed", err)
	}
	fmt.Println("mysql connect success!")
	selectPrepare, err := Db.Prepare("SELECT  id,account from wk_member where 1")
	fmt.Println(selectPrepare)

}
