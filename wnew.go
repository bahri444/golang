package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id_user  int
	username string
	password string
	level    string
}

func conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:bahri@tcp(127.0.0.1)/technology_code")
	if err != nil {
		return nil, err
	}
	return db, err
}
func GetUser() {
	db, err := conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	row, err := db.Query("select * from user")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer row.Close()
	var res []user
	for row.Next() {
		each := user{}
		err = row.Scan(&each.id_user, &each.username, &each.password, &each.level)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		res = append(res, each)
	}
	if row.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range res {
		fmt.Println(each.id_user, each.username, each.password, each.level)
	}
}
func main() {
	GetUser()
}
