package main

import "fmt"

func main() {
	var masuk, pulang, lama_kerja int
	fmt.Print("silahkan masukkkan jam masuk kerja :")
	fmt.Scanln(&masuk)
	fmt.Print("silahkan masukkan jam pulang kerja :")
	fmt.Scanln(&pulang)

	if pulang > masuk {
		lama_kerja = pulang - masuk
	} else {
		lama_kerja = (pulang + 12) - masuk
	}
	fmt.Println("anda kerja selama :", lama_kerja, " jam")
}
