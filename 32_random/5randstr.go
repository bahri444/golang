package main

import (
	"fmt"
	"math/rand"
	"time"
)

var data = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	rand.Seed(time.Now().UTC().UnixNano())                    //membuat angka random hasil eksekusi menjadi tidak pernah sama(tidak di ulang)
	fmt.Println("random string 5 karakter:", randomString(5)) //menampilkan random string ke terminal dengan index yg telah di tentukan
}

func randomString(panjang int) string {
	b := make([]rune, panjang)
	for i := range b {
		b[i] = data[rand.Intn(len(data))]
	}
	return string(b)
}
