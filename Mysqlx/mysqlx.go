package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID   int64
	Name sql.NullString
	Age  int
}

const (
	USERNAME   = ""
	PASSWORD   = ""
	NETWORK    = ""
	SERVERHOST = ""
	PORT       = ""
	DATABASE   = ""
)

func queryOne(DB *sql.DB) User {
	user := new(User)
	row := DB.QueryRow("select * from users where id=?", 1)
	if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed err:%v", err)
		return *user
	}
	return *user
}

/**
查询后不能不scan row ,否则会一直占用链接
 */
func quereyOneest(DB *sql.DB) User {
	user := new(User)
	for i := 0; i < 150; i++ {
		fmt.Println("query times:", i)
		row := DB.QueryRow("select * from users where id=? ", 1)
		continue
		if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			fmt.Println("can failed,err:%v", err)
			return *user
		}
		fmt.Println(*user)
	}
	return *user

}

func queryMulti(DB *sql.DB) map[int]User {
	u := map[int]User{}
	user := new(User)
	rows, err := DB.Query("select * from users where id>? order  by id desc ", 1)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return u
	}
	a := 0
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed,err:%v", err)
			return u
		}
		u[a] = *user
		a++
	}
	return u
}

func insertData(DB *sql.DB) {
	result, err := DB.Exec("insert  into users(name,age) values (?,?)", "huaer", 18)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsertId failed,err:%v", err)
		return
	}
	fmt.Println("LastInsertID", lastInsertID)
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RwosAffected", rowsaffected)

}
func updateData(DB *sql.DB) {
	result, err := DB.Exec("UPDATE users set age = ? where id=?", "30", 3)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

func deleteData(DB *sql.DB) {
	result, err := DB.Exec("delete from users where id=?", 1)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}
func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVERHOST, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if (err != nil) {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	fmt.Println("-------------------------------")
	user := queryOne(DB)
	fmt.Println("-------------------------------")
	fmt.Println(user.ID)
	//usera := quereyOneest(DB)
	//fmt.Println
	fmt.Println("-------------------------------")
	mu := queryMulti(DB)
	fmt.Println(mu)
	fmt.Println("-------------------------------")
	//顺序输出
	for a := 0; a < len(mu); a++ {
		fmt.Println(mu[a])
	}
	fmt.Println("-------------------------------")
	//非顺序输出
	for i := range mu {
		fmt.Println(mu[i])
	}
	fmt.Println("-------------------------------")
	insertData(DB)
	fmt.Println("-------------------------------")
	updateData(DB)
	fmt.Println("-------------------------------")
	deleteData(DB)
	fmt.Println("-------------------------------")

}
