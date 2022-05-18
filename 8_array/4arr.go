package main

import "fmt"

func main() {
	var buah = [...]string{"mangga", "jeruk", "kelapa", "durian", "semongko", "apple", "manggis", "kelengkeng", "nanas"}
	for i := 0; i < len(buah); i++ { //perulangan element array menggunakan for
		fmt.Printf("buahan %d : %s\n", i, buah[i])
	}
}
