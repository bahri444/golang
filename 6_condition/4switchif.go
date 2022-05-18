package main

import "fmt"

func main() {
	//switch case gaya if else
	var data = 3
	switch {
	case data == 9:
		fmt.Println("angka sembilan")
	case data == 8:
		fmt.Println("angka delapan")
	case data == 7:
		fmt.Println("angka tujuh")
	case (data < 6) && (data > 2):
		fmt.Println("angka kurang dari standar")
	case data == 0:
		fmt.Println("angka noll")
	}
}
