package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
	var luas = math.Pi * math.Pow(d/2, 2) //hitung lulas
	var keliling = math.Pi * d            //hitung keliling
	return luas, keliling                 //fungsi untuk mengembalikan 2 nilai(multiple return)
}
func main() {
	var diameter float64 = 18
	var luas, keliling = calculate(diameter)
	fmt.Printf("luas lingkaran :%.2f\n", luas)
	fmt.Printf("keliling lingkaran :%.2f\n", keliling)
}
