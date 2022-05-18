package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func dohashsalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltteks = fmt.Sprintf("text :%s\nsalt :%s", text, salt)
	fmt.Println(saltteks)
	var sha = sha1.New()
	sha.Write([]byte(saltteks))
	var enkripsi = sha.Sum(nil)
	return fmt.Sprintf("%x", enkripsi), salt
}
func main() {
	var text = "bahrysaepulajadeh"
	fmt.Printf("original :%s\n\n", text)
	var hashed1, salt1 = dohashsalt(text)
	fmt.Printf("hashed1 :%s\n\n", hashed1)
	var hashed2, salt2 = dohashsalt(text)
	fmt.Printf("hashed2 :%s\n\n", hashed2)
	var hashed3, salt3 = dohashsalt(text)
	fmt.Printf("hashed3 :%s\n\n", hashed3)
	_, _, _ = salt1, salt2, salt3
}
