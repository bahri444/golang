package main

import "fmt"

func main() {
	//closure adalah fungsi yang bisa di simpan dalam variabel
	//membuat fungsi di dalam fungsi atau bahkan membuat fungsi yang mengembalikan fungsi
	var minMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e > min:
				min = e
			}
		}
		return min, max
	}
	var nomer = []int{2, 3, 4, 5, 6, 8, 9, 7, 3}
	var min, max = minMax(nomer)
	fmt.Printf("data :%v\nmin :%v\nmax :%v\n", nomer, min, max)
}
