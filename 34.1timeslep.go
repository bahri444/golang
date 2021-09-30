package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 5) //cara satu
	fmt.Println("after 5 second")
	//////
	fmt.Println("Hello !")
	time.Sleep(2 * time.Second) //cara dua
}
