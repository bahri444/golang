package main

import (
	"flag"
	"fmt"
)

func main() {
	//cara pertama
	var data_nama = flag.String("nama", "saepul", "type your name")
	fmt.Println("nama :", *data_nama)

	//cara kedua
	var data_identitas string
	flag.StringVar(&data_identitas, "alamat", "desa selebung rembiga kecamatan janapria kabupaten lombok tengah", "type your identity")
	fmt.Println("alamat :", data_identitas)
}
