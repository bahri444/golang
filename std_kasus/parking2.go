package main

import "fmt"

func main() {
	var jenis_kendaraan, masuk, keluar, lama_parkir, bayar int
	fmt.Println("========================================")
	fmt.Println("| pilih jenis kendaraan anda ! [1/2].\t|")
	fmt.Println("|\t1. motor\t\t\t|")
	fmt.Println("|\t2. mobil\t\t\t|")
	fmt.Println("========================================")
	fmt.Scanln(&jenis_kendaraan)

	fmt.Print("jam masuk parkir :")
	fmt.Scanln(&masuk)
	fmt.Print("jam keluar parkir :")
	fmt.Scanln(&keluar)

	switch jenis_kendaraan {
	case 1:
		fmt.Println("jenis kendaraan anda adalah No", jenis_kendaraan, "motor")
	case 2:
		fmt.Println("jenis kendaraan anda adalah No", jenis_kendaraan, "motor")
	}
	if keluar > masuk {
		lama_parkir = keluar - masuk
	} else {
		lama_parkir = (keluar + 12) - masuk
	}
	for i := 1; i == jenis_kendaraan; i += 2 {
		bayar = lama_parkir*1000 + 4000
	}
	for i := 2; i == jenis_kendaraan; i += 2 {
		bayar = lama_parkir*1000 + 9000
	}
	fmt.Println("anda parkir selama :", lama_parkir, " jam")
	fmt.Println("maka anda harus membayar parkir senilai Rp.:", bayar)
}
