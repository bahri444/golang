package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) /*random int dengan cara ini
	outputnya tidak akan pernah sama(tidak akan pernah di ulang-ulang)
	walaupun di eksekusi dalam waktu	yang sama sehingga akan sangat
	sulit untuk di tebak angaka seed-nya(angka mulainya)*/
	fmt.Println("random ke-1", rand.Int())
	fmt.Println("random ke-2", rand.Int())
	fmt.Println("random ke-3", rand.Int())
	fmt.Println("random ke-4", rand.Int())
}
