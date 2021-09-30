package main

import (
	"fmt"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
func createFile() {
	path := "G:/fileCRUD/bonkon.txt"
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)
	if err == nil {
		fmt.Println("==> file sudah ada di dalam path")
	}
	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
			//if ini berfungsi jika path tidak ada maka dia akan mengembalikan ke fungsi isError yang di atas
		}
		fmt.Println("==> file berhasil dibuat", path)
		defer file.Close()
	}
}
func main() {
	createFile()
}
