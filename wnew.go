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
	alamat   string
	jen_kel  string
	telepon  string
	email    string
	akun_git string
	prodi    string
	kelas    string
	angkatan string
}

func Con() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:bahrysemet@tcp(127.0.0.1:3306)/ukm")
	if err != nil {
		return nil, err
	}
	return db, err
}
func querySql() {
	db, err := Con()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("insert into member values(?,?,?,?,?,?,?,?,?,?,?,?)", "11", "SI17200028", "septemi yuniati", "temi.jpg", "Perempuan", "waker puyung", "087757012337", "setemiYu@gmail.com", "yuniatiSept", "Sistem Informasi", "SI B", "pertama")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("data berhasil di insert")
	// defer rows.Close()
	// var res []member
	// for rows.Next() {
	// 	var each = member{}
	// 	err := rows.Scan(&each.id, &each.nim, &each.nama, &each.foto, &each.alamat, &each.kelas, &each.telepon, &each.email, &each.akun_git, &each.prodi, &each.kelas, &each.angkatan)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	res = append(res, each)
	// 	if rows.Err(); err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	for _, each := range res {
	// 		fmt.Println(each.id, each.nim, each.nama, each.foto, each.alamat, each.kelas, each.telepon, each.email, each.akun_git, each.prodi, each.kelas, each.angkatan)
	// 	}
	// }
}
func main() {
	querySql()
}
