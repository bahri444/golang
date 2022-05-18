package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "false"
	bulean, err := strconv.ParseBool(str)
	if err == nil {
		fmt.Println(bulean)
	}
	bul := true
	strenj := strconv.FormatBool(bul)
	fmt.Println(strenj)
}
