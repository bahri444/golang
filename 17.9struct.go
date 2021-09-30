package main

import "fmt"

type mhs struct {
	nama  string
	nilai int
}

func main() {
	//Deklarasi Anonymous Struct Menggunakan Keyword var
	var allmhs struct {
		nim int
		mhs
	}
	allmhs.mhs = mhs{"bahri", 99}
	allmhs.nim = 17200048
	fmt.Println(allmhs.nama)
	fmt.Println(allmhs.nilai)
	fmt.Println(allmhs.nim)
}
