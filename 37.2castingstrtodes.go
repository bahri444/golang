package main

import "fmt"

func main() {
	//casting string to decimal
	b := []byte("i love you")
	fmt.Printf("%d %d %d %d %d %d %d %d %d %d\n", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9])

	//casting decimal to string
	var str = []byte{105, 32, 108, 111, 118, 101, 32, 121, 111, 117}
	fmt.Printf("%s\n", str)
}
