package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//fungsi rand.Intn(n) untuk mendapatkan angka random pada indeks ke n.(indeks tertentu)
	fmt.Println("random int 3 at index :", rand.Intn(20))
}
