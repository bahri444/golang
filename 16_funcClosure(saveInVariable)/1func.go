package main

import (
	"fmt"
	"strings"
)

func search_data(data []string, call func(string) bool) []string {
	var result []string
	for _, each := range data {
		if searching := call(each); searching {
			result = append(result, each)
		}
	}
	return result
}
func main() {
	var data = []string{"saepul", "bahri", "emet", "bangepul"}
	var datacontainsU = search_data(data, func(each string) bool {
		return strings.Contains(each, "u")
	})
	var datalength6 = search_data(data, func(each string) bool {
		return len(each) == 6
	})
	fmt.Println("seluruh data yang ada : ", data)
	fmt.Println("data yang mengandung huruf U :", datacontainsU)
	fmt.Println("data yang mengandung enam karakter :", datalength6)
}
