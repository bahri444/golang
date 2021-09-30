package main

import "fmt"

func main() {
	var bilangan int = 294
	var bilangan2 *int = &bilangan

	fmt.Println("value pointer :", bilangan)
	fmt.Println("adress pointer:", &bilangan)

	fmt.Println("value pointer :", *bilangan2)
	fmt.Println("adress pointer:", bilangan2)

	bilangan = 245 //merubah nilai pada pointer
	fmt.Println("value pointer :", bilangan)
	fmt.Println("adress pointer:", &bilangan)

	fmt.Println("value pointer :", *bilangan2)
	fmt.Println("adress pointer:", bilangan2)
}
