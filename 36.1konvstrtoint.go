package main

import (
	"fmt"
	"strconv"
)

func main() {
	//strconv.Atoi()=>konversi data string ke integer
	str := "2340"
	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println("string to int :", num)
	}
	//strconv.Itoa()=>konversi integer ke string
	number := 2344
	stri := strconv.Itoa(number)
	fmt.Println("int to string :", stri)
}
