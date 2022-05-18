package main

import (
	"fmt"
	"os"
)

func removed() {
	var path = "G:/fileCRUD/bahry.txt"
	var err = os.Remove(path)
	if err == nil {
		fmt.Println("==> file berhasil di hapus")
	}
}
func main() {
	removed()
}
