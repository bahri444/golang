package main

import "fmt"

func main() {
	//tanda slice[]kurung siku kosong
	var data = []string{"nama", "nim", "alamat", "kelas"}
	fmt.Println("tampilkan data :", data[0:3]) //menampilkan element dari 0 sampai sebelum data ke 3
	fmt.Println("tampilkan data :", data[2])   //menampilkan elemen ke 3
	fmt.Println("tampilkan data :", data[2:4]) //menampilkan element ke 2 dan 3
}
