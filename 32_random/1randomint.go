package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)
	/*random standar jika di eksekusi kapanpun
	maka akan menghasilkan angaka random yang sama
	sehingga mudah untuk di tebak angka seed-nya(angka mulainya)*/
	fmt.Println("random ke-1", rand.Int())
	fmt.Println("random ke-2", rand.Int())
	fmt.Println("random ke-3", rand.Int())
	fmt.Println("random ke-4", rand.Int())
}
