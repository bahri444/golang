package main

import "fmt"

func main() {
	ordermakanan("kentucki")
	ordermakanan("piza")
}
func ordermakanan(menu string) {
	defer fmt.Println("terima kasih, silahkan di tunggu !")
	if menu == "piza" {
		fmt.Print("pilihan yang tepat !")
		fmt.Print("piza di tempat kami paling enak !")
		return
	}
	fmt.Println("pesanan anda :", menu)
}
