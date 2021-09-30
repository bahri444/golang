package main

import (
	"fmt"
	"regexp"
)

func main() {
	var data = "jambu strawbery nanas apel"
	var compileregex, _ = regexp.Compile(`[a-c]+`)
	var dataregex = compileregex.Split(data, -1) //-1 mengambil seluruh value yang ada dalam variabel data
	fmt.Printf("%#v", dataregex)
}
