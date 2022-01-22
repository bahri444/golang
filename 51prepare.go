package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type anggota_utc struct {
	no_agt    int
	nim       string
	nama      string
	foto      string
	jenis_kel string
	alamat    string
	email     string
	prodi     string
	kelas     string
	angkatan  string
}

func con() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mhs")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func sqlprepare() {
	db, err := con()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	var res = anggota_utc{}
	var no_agt = 1
	err = db.QueryRow("select nim,nama,foto,jenis_kel,alamat,email,prodi,kelas,angkatan from anggota_utc where no_agt=? ", no_agt).Scan(&res.nim, &res.nama, &res.foto, &res.jenis_kel, &res.alamat, &res.email, &res.prodi, &res.kelas, &res.angkatan)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s%s%s%s%s%s%s%s%s", res.nim, res.nama, res.foto, res.jenis_kel, res.alamat, res.email, res.prodi, res.kelas, res.angkatan)
}
func main() {
	sqlprepare()
}
