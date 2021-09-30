package main

import (
	"fmt"
	"strings"
)

func main() {
	var data interface{}
	data = 2
	//Casting Variabel Interface Kosong
	var sect = data.(int) * 10 //perlu casting ke tipe aslinya, yaitu int , setelahnya barulah nilai bisa dioperasikan, yaitu secret.(int) * 10 .
	fmt.Println(data, "kali sepuluh =", sect)

	data = []string{"jambu", "semongko", "jeruk", "manggis"}
	var fruit = strings.Join(data.([]string), ", ") //strings.Join=>untuk menggabungkan string yang satu dengan string yang lain
	fmt.Println("nama buah yang saya favoritkan :", fruit)
}
