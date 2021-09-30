package main

import "fmt"

//kombinasi deefer dan iife
func main() {
	number := 3
	if number == 3 {
		defer fmt.Println("hello 3")
		fmt.Println("hello 1")
	}
	fmt.Println("hello 2")
}
