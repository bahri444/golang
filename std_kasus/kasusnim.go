package main

import (
	"fmt"
)

func main() {
	var nim string
	fmt.Print("input nim : ")
	fmt.Scanln(&nim)
	var prodi string
	prodi = (nim[:2])
	if prodi == "SI" {
		fmt.Println(prodi, "(Sistem Informasi)")
		fmt.Println("NIM :", nim[2:])
	} else if prodi == "TI" {
		fmt.Println(prodi, "(Teknik Informatika)")
		fmt.Println("NIM :", nim[2:])
	} else {
		fmt.Println("NIM tidak terdaftar...!")
	}

}
