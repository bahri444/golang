package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type member struct {
	member_id  int
	nim        string
	nama       string
	tgl_lahir  string
	fas_foto   string
	kelamin    string
	alamat     string
	tlp        string
	email      string
	github     string
	prodi      string
	angkatan   int
	status_agt string
}

func con() (*sql.DB, error) {
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	db, err := sql.Open("mysql", "root:bahri@tcp(127.0.0.1:3306)/sis_techcode")
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
	var res = member{}
	var member_id = 1
	err = db.QueryRow("select nim,	nama,	tgl_lahir,	fas_foto,	kelamin,	alamat,	tlp,	email,	github,	prodi,	angkatan,	status_agt from member where member_id=? ", member_id).Scan(&res.nim, &res.nama, &res.tgl_lahir, &res.fas_foto, &res.kelamin, &res.alamat, &res.tlp, &res.email, &res.github, &res.prodi, &res.angkatan, &res.status_agt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%s%s%s%s%s%s%s%s%s%s%d%s\n", res.nim, res.nama, res.tgl_lahir, res.fas_foto, res.kelamin, res.alamat, res.tlp, res.email, res.github, res.prodi, res.angkatan, res.status_agt)

}
func main() {
	sqlprepare()
}
