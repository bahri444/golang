package main

import "fmt"

func main() {
	var buah = [...]string{"jeruk", "kelapa", "pisang", "jambu"}
	for i := range buah { //perulangan aray for-range
		fmt.Printf("buahan %d : %s\n", i, buah[i])
	}
}
