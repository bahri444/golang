package main

import "fmt"

func main() {
	//immediately invoked function expression
	//closure ini bisa di eksekusi langsung pada saat deklarasinya
	var nomer = []int{3, 4, 5, 6, 0, 4, 5, 6, 7, 8, 9}
	var newnomer = func(min int) []int {
		var r []int
		for _, e := range nomer {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)
	fmt.Println("nomer asli :", nomer)
	fmt.Println("filter nomer :", newnomer)
}
