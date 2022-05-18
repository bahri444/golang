package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (luas float64, keliling float64) {
	luas = math.Pi * math.Pow(d/2, 2)
	keliling = math.Pi * d
	return //fungsi tanpa nilai kembalian(fredefined return value)
}
func main() {
	var diameter float64 = 16
	var luas, keliling = calculate(diameter)
	fmt.Printf("luas lingkaran :%.2f\n", luas)
	fmt.Printf("keliling lingkaran :%2.f\n", keliling)
}
