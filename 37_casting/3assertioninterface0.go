package main

import "fmt"

func main() {
	var data = map[string]interface{}{
		"nama":   "Bahri",
		"alamat": "selebung",
		"nim":    17200048,
		"tinggi": 166.5,
		"berat":  65,
		"hoby":   []string{"rebahan", "dengerinmusic", "belajar teks warna warni"},
		"mhs":    true,
	}
	for _, value := range data {
		switch value.(type) {
		case string:
			fmt.Println(value.(string))
		case int:
			fmt.Println(value.(int))
		case bool:
			fmt.Println(value.(bool))
		case float64:
			fmt.Println(value.(float64))
		case []string:
			fmt.Println(value.([]string))
		default:
			fmt.Println("tipe ini belum ada penerimaannya :", value.(int))
		}
	}

}
