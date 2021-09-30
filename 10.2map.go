package main

import "fmt"

func main() {
	//iterasi map menggunakan for-range
	var keledai = map[string]int{
		"januari":  10,
		"februari": 12,
		"maret":    15,
		"april":    20,
	}
	for value, key := range keledai {
		fmt.Println(value, "\t:", key)
	}
}
