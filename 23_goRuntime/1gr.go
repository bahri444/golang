package main

import (
	"fmt"
	"runtime"
)

func Print(bil int, pesan string) {
	for i := 0; i < bil; i++ {
		fmt.Println((i + 1), pesan)

	}
}
func main() {
	runtime.GOMAXPROCS(2) //menentukan jumlah processor yang akan di pakai untuk eksekusi program
	go Print(3, "hallo")  //di eksekusi sbg go ruuntime baru
	Print(3, "apa kabar") //di eksekusi dengan cara biasa

	var input string
	fmt.Scanln(&input) //Fungsi fmt.Scanln() mengakibatkan proses jalannya aplikasi berhenti di baris itu
	//(blocking) hingga user menekan tombol enter.
}
