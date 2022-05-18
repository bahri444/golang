package main

import (
	"fmt"
	"runtime"
)

/*mirip seperti channel biasa.
Perbedaannya hanya pada penulisan deklarasinya, perlu ditambahkan angka
buffer sebagai argumen make()*/
func main() {
	runtime.GOMAXPROCS(2)
	pesan := make(chan int, 2)
	go func() {
		for {
			i := <-pesan
			fmt.Println("terima data ", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("send data :", i)
		pesan <- i
	}
}
