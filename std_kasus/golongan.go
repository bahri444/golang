package main

import "fmt"

func main() {
	var jamkerja, hasil int
	var golongan string
	fmt.Print("masukkan jam kerja :")
	fmt.Scanln(&jamkerja)
	fmt.Print("masukkan golongan :")
	fmt.Scanln(&golongan)
	switch golongan {
	case "A":
		if jamkerja <= 48 {
			hasil = jamkerja * 5000
			fmt.Println("gaji anda : ", hasil)
		} else if jamkerja >= 48 {
			hasil = (jamkerja-48)*3000 + (jamkerja * 5000)
			fmt.Println("gaji + bonus : ", hasil)
		}
		break
	case "B":
		if jamkerja <= 48 {
			hasil = jamkerja * 7000
			fmt.Println("gaji anda : ", hasil)
		} else if jamkerja >= 48 {
			hasil = (jamkerja-48)*5000 + (jamkerja * 7000)
			fmt.Println("gaji + bonus : ", hasil)
		}
		break
	case "C":
		if jamkerja <= 48 {
			hasil = jamkerja * 10000
			fmt.Println("gaji anda : ", hasil)
		} else if jamkerja >= 48 {
			hasil = (jamkerja-48)*7000 + (jamkerja * 10000)
			fmt.Println("gaji + bonus : ", hasil)
		}
		break

	}
}
