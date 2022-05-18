package main

import "fmt"

type mhs struct {
	nama  string
	nilai int
}

func main() {
	type people = mhs
	var p1 = people{"saepul", 45}
	fmt.Println(p1)

	var p2 = people{"bahri", 95}
	fmt.Println(p2)
}
