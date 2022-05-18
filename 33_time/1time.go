package main

import (
	"fmt"
	"time"
)

func main() {
	/*time.time merupakan representasi data-time untuk tipe tanggal dan waktu
	dua cara menentukan tanggal dan waktu yakni :
	1.time.Now()=> menentukan waktu secara otomatis
	2.time.Date()=> menentukan waktu secara manual (di tentukan sendiri oleh user)*/
	var timeone = time.Now()
	fmt.Printf("waktu sekarang :%v\n", timeone)

	var timetwo = time.Date(2021, 10, 01, 12, 0, 0, 0, time.UTC)
	fmt.Printf("time manualy :%v'n", timetwo)

}
