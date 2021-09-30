package main

import "fmt"

func main() {
	var masuk int
	var keluar int
	var total_jam int
	var bayar int
	fmt.Print("your input time login to parking :")
	fmt.Scanln(&masuk)
	fmt.Print("your input time logout into parking :")
	fmt.Scanln(&keluar)

	if keluar > masuk {
		total_jam = keluar - masuk
	} else {
		total_jam = (keluar + 12) - masuk
	}
	for i := 0; i < total_jam; i += 2 {
		bayar = total_jam*1000 + 4000
	}
	fmt.Println("anda parkir selama :", total_jam, " jam")
	fmt.Println("maka anda harus membayar parkir sebanyak Rp.", bayar)
}
