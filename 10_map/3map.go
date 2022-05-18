package main

import "fmt"

func main() {
	var data = map[string]int{"september": 89, "desember": 100, "januari": 20, "april": 150}
	fmt.Println(len(data), data)

	delete(data, "januari") //delete perintah untuk menghapus data
	fmt.Println(len(data), data)

}
