package main

import "fmt"

type mhs struct {
	nama  string
	nilai int
}

func main() {
	var sis = mhs{nama: "saepul", nilai: 99}
	var sis2 *mhs = &sis
	fmt.Println("nama1 :", sis.nama, "\tnilai :", sis.nilai)
	fmt.Println("nama2 :", sis2.nama, "\tnilai :", sis.nilai)

	sis2.nama = "bahri"
	fmt.Println("nama1 :", sis.nama, "\tnilai :", sis.nilai)
	fmt.Println("nama2 :", sis2.nama, "\tnilai :", sis.nilai)
}
