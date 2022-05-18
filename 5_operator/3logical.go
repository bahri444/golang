package main

import "fmt"

func main() {
	var kiri = true
	var kanan = false
	var kiridankanan = kiri && kanan
	fmt.Println(kiridankanan)

	kiridankanan = kiri || kanan
	fmt.Println(kiridankanan)

	var notkanan = !kiri
	fmt.Println(notkanan)
}
