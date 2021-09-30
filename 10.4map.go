package main

import "fmt"

func main() {
	var data = map[string]int{"id": 1, "nim": 17200048}
	var value, esist = data["kelas"] //mendeteksi keberadaan data(item) dengan key tertentu
	if esist {
		fmt.Println(value, "data validion")
	} else {
		fmt.Println("data is not validation")
	}
}
