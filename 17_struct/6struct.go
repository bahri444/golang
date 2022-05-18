package main

import "fmt"

type mhs struct {
	nama string
	nim  int
}

func main() {
	// Anonymous struct adalah struct yang tidak dideklarasikan di awal sebagai tipe data baru, melainkan langsung ketika pembuatan objek
	var sis = struct {
		nilai int
		mhs
	}{}
	/* anonymous struct dengan pengisian property
	var s2 = struct {
	person
	grade int
	}{
	person: person{"wick", 21},
	grade: 2,*/
	sis.mhs = mhs{"saepul bahri", 17200048}
	sis.nilai = 100
	fmt.Println("nama :", sis.mhs.nama)
	fmt.Println("nim  :", sis.mhs.nim)
	fmt.Println("nilai:", sis.nilai)

}
