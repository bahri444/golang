package main

import (
	"fmt"
	"regexp"
)

func main() {
	var date = "jambu srawbery nanas apel"
	var regex, err = regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}
	var res = regex.FindAllString(date, 2) //mengambil dua elemen string
	fmt.Printf("%#v\n", res)
	var res2 = regex.FindAllString(date, -1) //mengambil semua elemen string yang ada
	fmt.Printf("%#v\n", res2)
	var res3 = regex.FindAllString(date, 1) //mengambil satu elemen string
	fmt.Printf("%#v\n", res3)
}
