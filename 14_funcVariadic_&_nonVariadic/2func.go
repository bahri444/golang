package main

import "fmt"

//closure(funsi di dalam fungsi,bahkan fungsi mengembalikan fungsi)sebagai nilai kemblian
func finmax(number []int, max int) (int, func() []int) {
	var res []int
	for _, e := range number {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
func main() {
	var max = 3
	var number = []int{2, 3, 4, 5, 3, 5, 6, 7, 8, 0}
	var how, getnum = finmax(number, max)
	var thenum = getnum()

	fmt.Println("numbers :", number)
	fmt.Printf("maxs \t: %d\n\n", max)
	fmt.Println("found \t:", how)
	fmt.Println("value \t:", thenum)
}
