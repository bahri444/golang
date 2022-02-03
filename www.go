package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("masukkan nilai : ")
	fmt.Scan(&nilai)
	if nilai >= 75 {
		if nilai == 100 {
			fmt.Println("nilai sangat bagus")
		} else if nilai >= 90 {
			fmt.Println("nilai bagus")
		} else if nilai >= 80 {
			fmt.Println("nilai B")
		} else if nilai >= 75 {
			fmt.Println("nilai standar")
		}
	} else {
		if nilai >= 70 {
			fmt.Println("nilai C")
		} else if nilai >= 60 {
			fmt.Println("nilai C-")
		} else if nilai >= 50 {
			fmt.Println("nilai D")
		} else if nilai <= 40 {
			fmt.Println("nilai E")
		}
	}
}
