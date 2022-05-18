package main

import "fmt"

func main() {
	//Kombinasi Slice, map , dan interface{}
	var data = []map[string]interface{}{
		{"nama": "bahri", "nilai": 80},
		{"nama": "saepul", "nilai": 89},
		{"nama": "saputra", "nilai": 85},
	}
	for _, each := range data {
		fmt.Println(each["nama"], ": nilai is :", each["nilai"])
	}
	//cara deklarasi ke 2
	fmt.Println()
	var buahan = []interface{}{
		map[string]interface{}{"nama": "kelapa", "total": 10},
		[]string{"mangga", "jeruk", "kelapa", "manggis"},
		"jeruti",
	}
	for _, each := range buahan {
		fmt.Println(each)
	}
}
