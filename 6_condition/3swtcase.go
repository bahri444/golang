package main

import "fmt"

func main() {
	var bil = 9
	switch bil {
	case 9:
		fmt.Println("angka sembilan")
	case 8:
		fmt.Println("angka delapan")
	case 7:
		fmt.Println("angka tujuh")
	case 6, 5, 4: //contoh penggunaan switch case dalam banyak kondisi
		fmt.Println("angka enam")
	default: //contoh penggunaan switch case dan default
		fmt.Println("angka tidak valid")
	}
}
