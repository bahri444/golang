package main

import "fmt"

func main() {
	//pemanfaatan label dalam perulangan
outerloop:
	for i := 0; i < 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Print("matriks :[", i, "][", j, "]", "\n")
		}
	}
}
