package main

import "fmt"

func main() {
	var bil1 int = 200
	var bil2 *int = &bil1
	fmt.Println("value bilangan1 :", bil1)
	fmt.Println("adres bilangan1 :", &bil1)

	fmt.Println("value bilangan2 :", *bil2)
	fmt.Println("adres bilangan2 :", bil2)
}
