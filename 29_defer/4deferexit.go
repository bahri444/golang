package main

import "fmt"

//kombinasi defer dan iife bagian dua
func main() {
	bil := 3
	if bil == 3 {
		fmt.Println("hello 1")
		func() {
			defer fmt.Println("hello 3")
		}()
	}
	fmt.Println("hello 2")
	//Agar halo 3 bisa dimunculkan di akhir blok if, maka harus dibungkus dengan IIFE.
}
