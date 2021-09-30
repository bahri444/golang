package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	var teks = "contoh enkripsi kode"
	var sha = sha1.New()
	sha.Write([]byte(teks))
	var enkripsi = sha.Sum(nil)
	var enkripsistr = fmt.Sprintf("%x", enkripsi)
	fmt.Println(enkripsistr)
}
