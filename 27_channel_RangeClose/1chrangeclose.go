package main

import (
	"fmt"
	"runtime"
)

func sharepesan(ch chan<- string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("bilangan :%d\n", i)
	}
	close(ch) //Fungsi close digunakan utuk me-non-aktifkan channel.
}
func printpesan(ch <-chan string) {
	for pesan := range ch { //Channel yang sudah di-close tidak bisa digunakan lagi untuk menerima dan mengirim data for-range berhenti juga
		fmt.Println(pesan)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	var pesan = make(chan string)
	go sharepesan(pesan)
	printpesan(pesan)
	/*channel direction
	ch chan string = kirim dan terima data
	ch chan<- string=kirim data
	ch <-chan string=terima data*/
}
