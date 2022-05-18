package main

import "fmt"

func main() {
	//for tanpa argumen
	var i = 1
	for {
		fmt.Println("bilangan bulat", i)
		i++
		if i == 5 {
			break
		}
	}
}
