package main

import "fmt"

func main() {
	var bil1 = [2][3]int{{9, 8, 7}, {1, 2, 3}}
	var bil2 = [2][3]int{[3]int{3, 2, 1}, [3]int{6, 5, 4}}
	fmt.Println("bilangan 1:", bil1)
	fmt.Println("bilangan 2:", bil2)
}
