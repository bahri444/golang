package main

import "fmt"

type mhs struct {
	nama string
	nim  int
}

// Embedded Struct Dengan Nama Property Yang Sama
type mhi struct {
	nim   int
	nilai int
	mhs
}

func main() {
	var sis = mhi{}
	sis.nama = "saepul bahri"
	sis.nim = 17200048     //versi mhi
	sis.mhs.nim = 17200050 //versi mhs
	sis.nilai = 99

	fmt.Println("nama :", sis.nama)
	fmt.Println("nim  :", sis.nim)
	fmt.Println("nim  :", sis.mhs.nim)
	fmt.Println("nilai:", sis.nilai)
}
