package main

import "fmt"

func main() {
	var data = []map[string]string{
		map[string]string{"nama": "saepul bahri", "alamat": "melar desa selebung"},
		map[string]string{"nama": "feerdi barliansyah", "alamat": "kopang desa kopang"},
		map[string]string{"nama": "lalu ata sadli", "alamat": "muncan desa muncan"},
	}
	for _, identitas := range data {
		fmt.Println(identitas["nama"], identitas["alamat"])
	}
}
