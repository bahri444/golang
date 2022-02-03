package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type member struct {
	Id       string
	Nim      string
	Nama     string
	Foto     string
	Jen_kel  string
	Alamat   string
	Telepon  string
	Email    string
	Akun_git string
	Prodi    string
	Kelas    string
	Angkatan string
}
type Responses struct {
	Status bool
	Pesan  string
	Data   []member
}

func Koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ukm")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func Read(Pesane string) Responses {
	db, err := Koneksi()
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "gagal koneksi ke database" + err.Error(),
			Data:   []member{},
		}
	}
	defer db.Close()
	rows, err := db.Query("select * from member order by id asc")
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "Query salah" + err.Error(),
			Data:   []member{},
		}
	}
	defer rows.Close()
	var res []member
	for rows.Next() {
		each := member{}
		err = rows.Scan(&each.Id, &each.Nim, &each.Nama, &each.Foto, &each.Jen_kel, &each.Alamat, &each.Telepon, &each.Email, &each.Akun_git, &each.Prodi, &each.Kelas, &each.Angkatan)
		if err != nil {
			return Responses{
				Status: false,
				Pesan:  "nama rows ada yang tidak di temukan" + err.Error(),
				Data:   []member{},
			}
		}
		res = append(res, each)
	}
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "data gagal di tampilkan" + err.Error(),
			Data:   []member{},
		}
	}
	return Responses{
		Status: true,
		Pesan:  Pesane,
		Data:   res,
	}
}
func GetMember(Id string) Responses {
	db, err := Koneksi()
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "gagal koneksi ke database" + err.Error(),
			Data:   []member{},
		}
	}
	defer db.Close()
	rows, err := db.Query("select * from member  where id=?", Id)
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "Query salah" + err.Error(),
			Data:   []member{},
		}
	}
	defer rows.Close()
	var res []member
	for rows.Next() {
		each := member{}
		err = rows.Scan(&each.Id, &each.Nim, &each.Nama, &each.Foto, &each.Jen_kel, &each.Alamat, &each.Telepon, &each.Email, &each.Akun_git, &each.Prodi, &each.Kelas, &each.Angkatan)
		if err != nil {
			return Responses{
				Status: false,
				Pesan:  "row ada yang tidak di temukan" + err.Error(),
				Data:   []member{},
			}
		}
		res = append(res, each)
	}
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "gagal tampil berdasarkan id" + err.Error(),
			Data:   []member{},
		}
	}
	return Responses{
		Status: true,
		Pesan:  "yes berhasil di tampilkan",
		Data:   res,
	}
}
func Insert(id string, nim string, nama string, foto string, jen_kel string, alamat string, telepon string, email, string, akun_git string, prodi string, kelas string, angkatan string) Responses {
	db, err := Koneksi()
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "gagal connection in func insert" + err.Error(),
			Data:   []member{},
		}
	}
	defer db.Close()
	_, err = db.Exec("insert into member values(?,?,?,?,?,?,?,?,?,?,?,?)", id, nim, nama, foto, jen_kel, alamat, telepon, email, akun_git, prodi, kelas, angkatan)
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "data gagal di tambahkan" + err.Error(),
			Data:   []member{},
		}
	}
	return Responses{
		Status: true,
		Pesan:  "data berhasil di tambahkan",
		Data:   []member{},
	}
}
func Update(id string, nim string, nama string, foto string, jen_kel string, alamat string, telepon string, email, string, akun_git string, prodi string, kelas string, angkatan string) Responses {
	db, err := Koneksi()
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "gagal connection in func update" + err.Error(),
			Data:   []member{},
		}
	}
	defer db.Close()
	_, err = db.Exec("UPDATE member SET nim=?,nama=?,foto=?,jenkel=?,alamat=?,telepon=?,email=?,akun_git=?,prodi=?,kelas=?,angkatan=? WHERE id=?", nim, nama, foto, jen_kel, alamat, telepon, email, akun_git, prodi, kelas, angkatan)
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "data gagal di edit" + err.Error(),
			Data:   []member{},
		}
	}
	return Responses{
		Status: true,
		Pesan:  "data berhasil di edit" + err.Error(),
		Data:   []member{},
	}
}
func Delete(id string) Responses {
	db, err := Koneksi()
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "gagal connection in func delete" + err.Error(),
			Data:   []member{},
		}
	}
	defer db.Close()
	_, err = db.Exec("delete from member where id=?")
	if err != nil {
		return Responses{
			Status: false,
			Pesan:  "data gagal di hapus" + err.Error(),
			Data:   []member{},
		}
	}
	return Responses{
		Status: true,
		Pesan:  "data berhasil di hapus",
		Data:   []member{},
	}
}
func Controller(w http.ResponseWriter, r *http.Request) {
	ReadHtml, errHtml := template.ParseFiles("template/Read.html")
	if errHtml != nil {
		fmt.Println(errHtml.Error())
		return
	}
	InsertHtml, errHtml := template.ParseFiles("template/Insert.html")
	if errHtml != nil {
		fmt.Println(errHtml.Error())
		return
	}
	UpdateHtml, errHtml := template.ParseFiles("template/Update.html")
	if errHtml != nil {
		fmt.Println(errHtml.Error())
		return
	}
	DeleteHtml, errHtml := template.ParseFiles("template/Delete.html")
	if errHtml != nil {
		fmt.Println(errHtml.Error())
		return
	}
	switch r.Method {
	case "GET":
		aksi := r.URL.Query()["aksi"]
		if len(aksi) == 0 {
			ReadHtml.Execute(w, Read("data berhasil di load"))
		} else if aksi[0] == "Insert" {
			InsertHtml.Execute(w, nil)
		} else if aksi[0] == "Update" {
			id := r.URL.Query()["id"]
			UpdateHtml.Execute(w, GetMember(id[0]))
		} else if aksi[0] == "Delete" {
			id := r.URL.Query()["id"]
			DeleteHtml.Execute(w, GetMember(id[0]))
		} else {
			ReadHtml.Execute(w, "data aja deh")
		}
	case "POST":
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "gagal", err)
			return
		}
		id := r.FormValue("id")
		nim := r.FormValue("nim")
		nama := r.FormValue("nama")
		foto := r.FormValue("foto")
		jen_kel := r.FormValue("jen_kel")
		alamat := r.FormValue("alamat")
		telepon := r.FormValue("telepon")
		email := r.FormValue("email")
		akun_git := r.FormValue("akun_git")
		prodi := r.FormValue("prodi")
		kelas := r.FormValue("kelas")
		angkatan := r.FormValue("angkatan")
		aksi := r.URL.Path
		if aksi == "/Insert" {
			res := Insert(id, nim, nama, foto, jen_kel, alamat, telepon, email, akun_git, prodi, kelas, angkatan, "")
			ReadHtml.Execute(w, Read(res.Pesan))
		} else if aksi == "/Update" {
			res := Update(id, nim, nama, foto, jen_kel, alamat, telepon, email, akun_git, prodi, kelas, angkatan, "")
			ReadHtml.Execute(w, Read(res.Pesan))
		} else if aksi == "/Delete" {
			res := Delete(id)
			ReadHtml.Execute(w, Read(res.Pesan))
		} else {
			ReadHtml.Execute(w, Read("data berhasil di tampilkan"))
		}
	default:
		fmt.Fprintln(w, "method tidak tersedia")
	}
}
func main() {
	http.HandleFunc("/", Controller)
	fmt.Println("server aktivate on port 8081")
	http.ListenAndServe(":8081", nil)
}
