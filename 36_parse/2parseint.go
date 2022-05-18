package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "200000000000"
	numer, err := strconv.ParseInt(str, 10, 64) //64=tipe data int64
	if err == nil {
		fmt.Println(numer)
	}
	strbin := "1010"
	bilangan, err := strconv.ParseInt(strbin, 2, 8) //8=tipe data int8 // angka 2 untuk mengambil dua digit angka biner
	if err == nil {
		fmt.Println(bilangan)
	}
}
