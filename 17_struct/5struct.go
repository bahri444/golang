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

func main() {
	// Pengisian nilai property sub-struct bisa dilakukan dengan langsung memasukkan variabel objek yang tercetak dari struct yang sama.
	var sis = mhs{nama: "bahri", nim: 17200048}
	var sem = mhi{mhs: sis, nilai: 89}
	fmt.Println("nama :", sem.nama)
	fmt.Println("nama :", sem.nim)
	fmt.Println("nama :", sem.nilai)
}
