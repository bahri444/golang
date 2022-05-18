package main

import "fmt"

func main() {
	var nilai = (((2 * 10) + 20%3) - 10) / 2
	var eksekusi = (nilai == 2)
	fmt.Printf("hasil %d (%t) \n", nilai, eksekusi)
}
