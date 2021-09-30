package main

import "fmt"

func change(original *int, value int) {
	*original = value //manipulasi parameter pointer
}
func main() {
	var bilangan = 203
	fmt.Println("sebelum :", bilangan)

	change(&bilangan, 10)
	fmt.Println("setelah :", bilangan)
}
