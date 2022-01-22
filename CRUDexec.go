package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type login struct {
	id        int
	username  string
	password  string
	password2 string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mhs")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("insert into login values(?,?,?,?)", "09", "fileadmin", "admin", "admin")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("data berhasil di input")
}
func main() {
	sqlExec()
}
