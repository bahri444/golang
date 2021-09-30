package main

import "fmt"

func main() {
	//slice merupakan tipe data refrensi
	var buahan = []string{"kelapa", "mangga", "jeruk", "jambu"}

	var buahan1 = buahan[0:3]
	var buahan2 = buahan[1:4]

	var buahan3 = buahan1[1:2]
	var buahan4 = buahan2[0:1]

	fmt.Println(buahan)
	fmt.Println(buahan1)
	fmt.Println(buahan2)
	fmt.Println(buahan3)
	fmt.Println(buahan4)

	//elemen ke 0 akan di rubah menjadi semongko
	buahan4[0] = "semongko"
	fmt.Println(buahan)
	fmt.Println(buahan1)
	fmt.Println(buahan2)
	fmt.Println(buahan3)
	fmt.Println(buahan4)
	//9.2fungsi len pada slice
	fmt.Println("jumlah seluruh element slice ", len(buahan))
}
