package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type anggota struct {
	id            int
	foto          string
	nim           string
	nama          string
	jenis_kelamin string
	alamat        string
	email         string
	prodi         string
	kelas         string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mahasiswa")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from anggota")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []anggota

	for rows.Next() {
		var each = anggota{}
		var err = rows.Scan(&each.id, &each.foto, &each.nim, &each.nama, &each.jenis_kelamin, &each.alamat, &each.email, &each.prodi, &each.kelas)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.id, each.foto, each.nim, each.nama, each.jenis_kelamin, each.alamat, each.email, each.prodi, each.kelas)
	}
}

func main() {
	sqlQuery()
}
