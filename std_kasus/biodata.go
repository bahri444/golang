package main

import (
	"fmt"
	"time"
)

func main() {
	var nim string
	var nama string
	var alamat string
	var status string
	var tahun int
	var th = time.Now()
	fmt.Print("input nim :")
	fmt.Scanln(&nim)
	fmt.Print("input nama :")
	fmt.Scanln(&nama)
	fmt.Print("input tahun lahir :")
	fmt.Scanln(&tahun)
	fmt.Print("input alamat :")
	fmt.Scanln(&alamat)
	fmt.Print("input status :")
	fmt.Scanln(&status)
	fmt.Println("\tBIODATA")
	var prodi = (nim[:2])
	var nimm = (nim[2:])
	if prodi == "TI" {
		fmt.Println("Prodi \t:", prodi, "(Teknik Informatika)")
	} else if prodi == "SI" {
		fmt.Println("Prodi \t:", prodi, "(Sistem Informasi)")
	}
	fmt.Println("Nim \t:", nimm)
	fmt.Println("Nama \t:", nama)
	fmt.Println("Usia \t:", th.Year()-tahun, "tahun")
	fmt.Println("Alamat \t:", alamat)
	fmt.Println("Status \t:", status)

}
