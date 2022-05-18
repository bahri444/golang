package main

import "fmt"

func main() {
	//pembuatan parameter dengan fungsi sejenis(variadic) yang tak terbatas
	var bil = calculate(4, 4, 3, 5, 6, 7, 8, 9, 9)
	var tampung = fmt.Sprintf("rata-rata :%2.f", bil)
	fmt.Println(tampung)
}
func calculate(number ...int) float64 {
	var total int = 0
	for _, no := range number {
		total += no
	}
	var bil = float64(total) / float64(len(number))
	return bil
}
