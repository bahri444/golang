package main

import (
	"fmt"
	"regexp"
)

func main() {
	var data = "BANUN WHY123 semongko"
	var compileregex, _ = regexp.Compile(`[a-z]+`)
	var dataregex = compileregex.FindString(data) //deteksi value string jika semua string maka di kembalikan value yang paling kiri
	fmt.Println(dataregex)
}
