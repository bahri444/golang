package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*membankitkan angka random dalam tipe uint32 dan float32*/
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random int :", rand.Int())
	fmt.Println("random float32 :", rand.Float32())
	fmt.Println("random uint :", rand.Uint32())
}
