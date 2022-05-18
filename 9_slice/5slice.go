package main

import "fmt"

func main() {
	var buah = []string{"jambu", "semangka", "lemon", "pepaya"}
	var buah2 = buah[0:3]
	var buah3 = buah[0:3:3] //akses slice menggunakan 3 indexs
	fmt.Println(buah)
	fmt.Println(len(buah))
	fmt.Println(cap(buah))

	fmt.Println(buah2)
	fmt.Println(len(buah2))
	fmt.Println(cap(buah2))

	fmt.Println(buah3)
	fmt.Println(len(buah3))
	fmt.Println(cap(buah3))
}
