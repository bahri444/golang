package main

import "fmt"

func main() {
	var buahan = []string{"jambu", "kelapa", "strawbery"}
	var abuahan = buahan[0:2]
	fmt.Println(len(abuahan))
	fmt.Println(cap(abuahan))

	fmt.Println(buahan)
	fmt.Println(abuahan)

	var fruit = append(buahan, "markisa") //append untuk menambah element pada slice
	fmt.Println(buahan)
	fmt.Println(abuahan)
	fmt.Println(fruit)
}
