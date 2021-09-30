package main

import "fmt"

func main() {
	//perulangan bersarang
	for i := 1; i <= 5; i++ {
		for j := i; j <= 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
