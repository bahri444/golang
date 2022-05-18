package main

import "fmt"

func main() {
	var buahan = [...]string{"jeruk", "kelapa", "mangga", "jambu"}
	for _, buah := range buahan { //perulangan array for-range menggunakan variabel underscore
		fmt.Printf("buahan : %s\n", buah)
	}
}
