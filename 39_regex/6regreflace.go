package main

import (
	"fmt"
	"regexp"
)

func main() {
	var data = "jambu strawbery nanas apel"
	var compileregex, _ = regexp.Compile(`[a-z]+`)
	var dataregex = compileregex.ReplaceAllStringFunc(data, func(each string) string {
		if each == "nanas" {
			return "pisang" //menganti nanas dengan pisang
		}
		return each
	})
	fmt.Println(dataregex)
}
