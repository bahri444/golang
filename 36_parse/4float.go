package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "24.12"
	num, err := strconv.ParseFloat(str, 32)
	if err == nil {
		fmt.Println(num)
	}
}
