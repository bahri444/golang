package main

import (
	"fmt"
	"strings"
)

//fungsi dengan parameter biasa dengan parameter variadic(sejenis)
func myhoby(nama string, hobi ...string) {
	var hobiAsStr = strings.Join(hobi, ", ")

	fmt.Printf("hello, nama saya : %s\n", nama)
	fmt.Printf("my hobi is  :%s\n", hobiAsStr)
}
func main() {
	myhoby("bahry", "ngoding", "dengerin musik", "rebahan")
}
