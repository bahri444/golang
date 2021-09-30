package main

import "fmt"

type mhs struct {
	nama  string
	nilai int
}

func main() {
	var sis = mhs{} //cara inisialisasi objek struct
	sis.nama = "bahri"
	sis.nilai = 90

	var sis2 = mhs{"saepul", 89}            //cara deklarasi ke2
	var sis3 = mhs{nama: "emet", nilai: 99} //cara deklarasi ke3 atau bisa di deklarasikan nilainya secara acak

	fmt.Println("nama :", sis.nama, "\tnilai :", sis.nilai)
	fmt.Println("nama :", sis2.nama, "\tnilai :", sis2.nilai)
	fmt.Println("nama :", sis3.nama, "\tnilai :", sis3.nilai)
}
