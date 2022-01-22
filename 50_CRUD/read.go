package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type member struct {
	id       int
	nim      string
	nama     string
	foto     string
	jen_k    string
	alamat   string
	telepon  string
	email    string
	akun_git string
	prodi    string
	kelas    string
	angkatan string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:bahrysemet@tcp(127.0.0.1:3306)/ukm")
	if err != nil {
		return nil, err
	}
	return db, err
}
func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from member")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var res []member
	for rows.Next() {
		var each = member{}
		err := rows.Scan(&each.id, &each.nim, &each.nama, &each.foto, &each.jen_k, &each.alamat, &each.telepon, &each.email, &each.akun_git, &each.prodi, &each.kelas, &each.angkatan)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		res = append(res, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range res {
		fmt.Println(each.id, each.nim, each.nama, each.foto, each.jen_k, each.alamat, each.telepon, each.email, each.akun_git, each.prodi, each.kelas, each.angkatan)
	}
}
func main() {
	sqlQuery()
}
