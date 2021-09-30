package main

import (
	"fmt"
	"regexp"
)

func main() {
	var date = "1234 strawbery 648236 9897"
	var compileregex, _ = regexp.Compile(`[a-z]+`)
	var dataregex = compileregex.FindStringIndex(date) //mengambil elemen index pada string
	fmt.Println(dataregex)
	fmt.Println(date[5:14])
}
