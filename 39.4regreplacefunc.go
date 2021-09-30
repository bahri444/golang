package main

import (
	"fmt"
	"regexp"
)

func main() {
	var data = "jambu strawbey nanas apel"
	var compileregex, _ = regexp.Compile(`[a-z]+`)
	var dataregex = compileregex.ReplaceAllString(data, "nanas") //mengulang sebuah kallimat tertentu sampai sebanyak jumlah kalimat yang ada.
	fmt.Println(dataregex)
}
