package main

import (
	"fmt"
	"runtime"
)

//Eksekusi Goroutine Pada IIFE
func main() {
	runtime.GOMAXPROCS(2)
	var pesan = make(chan string)
	go func(who string) {
		var data = fmt.Sprintf("hello :%s", who)
		pesan <- data
	}("bahri")
	var pesanan = <-pesan
	fmt.Println(pesanan)
}
