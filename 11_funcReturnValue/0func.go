package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = []string{"saepul", "bahri"}
	cetak("hello", name)
}
func cetak(name string, arr []string) {
	var namestr = strings.Join(arr, " ")
	fmt.Println(name, namestr)
}
