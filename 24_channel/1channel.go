package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var pesan = make(chan string) //deklarasi channel

	var helloto = func(who string) { // sayHelloTo adalah closure untuk membuat sebuah
		var data = fmt.Sprintf("hello :%s", who) //pesan string yang kemudian dikirim via channel. Pesan string tersebut dikirim
		pesan <- data                            //lewat channel messages .
	}
	//Fungsi sayHelloTo dieksekusi tiga kali sebagai goroutine berbeda. Menjadikan
	//tiga proses ini berjalan secara asynchronous atau tidak saling tunggu.
	go helloto("saepul emet")
	go helloto("bang emet")
	go helloto("saputra bahri")

	//proses pengiriman dan penerimaan data dari varaiabel pesan ke variabel pesan 1,2,3
	var pesan1 = <-pesan
	fmt.Println(pesan1)
	var pesan2 = <-pesan
	fmt.Println(pesan2)
	var pesan3 = <-pesan
	fmt.Println(pesan3)

}
