package main

import "fmt"

func main() {
	var data interface{}
	data = "saepul bahri"
	fmt.Println("data 1:", data)

	data = []string{"mangga", "kelapa", "jeruk", "kelengkeng"}
	fmt.Println("data 2:", data)

	data = 12.99
	fmt.Println("data 3:", data)
}
