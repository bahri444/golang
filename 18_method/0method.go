package main

import (
	"fmt"
	"strings"
)

type mhs struct {
	nama  string
	nilai int
}

func (m mhs) hello() {
	fmt.Println("hello ! ", m.nama)
}
func (m mhs) getname(i int) string {
	return strings.Split(m.nama, " ")[i-1]
}
func main() {
	var m1 = mhs{"saepul bahri", 90}
	m1.hello()
	var nama = m1.getname(2)
	fmt.Println("nick name :", nama)
}
