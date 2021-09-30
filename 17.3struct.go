package main

import "fmt"

type mhs struct {
	nama string
	nim  int
}
type mhi struct {
	nilai int
	mhs
}

//Embedded struct => mekanisme untuk menempelkan sebuah struct sebagai properti struct lain.
func main() {
	var sis = mhi{}
	sis.nama = "saepul bahri"
	sis.nim = 1720048
	sis.nilai = 99
	fmt.Println("nama :", sis.nama)
	fmt.Println("nim  :", sis.nim)
	fmt.Println("nilai:", sis.mhs.nim)
	fmt.Println("nilai:", sis.nilai)
}
