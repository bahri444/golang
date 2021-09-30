package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = int64(14)
	var bil = strconv.FormatInt(str, 8)
	fmt.Println(bil)

}
