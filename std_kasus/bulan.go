package main

import "fmt"

func main() {
	var pilihan int
	fmt.Print("masukkan angka :")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		fmt.Println("january")
		break
	case 2:
		fmt.Println("february")
		break
	case 3:
		fmt.Println("maret")
		break
	case 4:
		fmt.Println("april")
		break
	case 5:
		fmt.Println("mei")
		break
	case 6:
		fmt.Println("juni")
		break
	case 7:
		fmt.Println("juli")
		break
	case 8:
		fmt.Println("agustus")
		break
	case 9:
		fmt.Println("september")
		break
	case 10:
		fmt.Println("oktober")
		break
	case 11:
		fmt.Println("november")
		break
	case 12:
		fmt.Println("desember")
		break

	}
}
