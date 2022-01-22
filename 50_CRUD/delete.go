package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type anggota struct {
	no_agt   int
	nim      string
	nama     string
	foto     string
	jen_k    string
	alamat   string
	email    string
	prodi    string
	kelas    string
	angkatan string
}

func conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Bahry321@@tcp(127.0.0.1:3306)/ukm")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func sqlExec() {
	db, err := conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	_, err = db.Exec("delete from anggota where no_agt=?", 12)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("data berhasil di hapus")
}
func main() {
	sqlExec()
}
