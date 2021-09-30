package main

import (
	"fmt"
	"regexp"
)

func main() {
	var nama = "jambu strawbery nanas apel"
	var regex, _ = regexp.Compile(`[a-z]+`)
	var match = regex.MatchString(nama) //mendeteksi nilai string apakah memenuhi kriteria regex
	fmt.Println(match)
}
