package main

import "fmt"

type mhs struct {
	nama string
	nim  int
} //cara deklarasi struct
//penerapan struct
func main() {
	var sis mhs
	sis.nama = "saepul bahri"
	sis.nim = 17200048

	fmt.Println("nama mahasiswa :", sis.nama)
	fmt.Println("nim mahasiswa  :", sis.nim)
}
