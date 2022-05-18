package main

import (
	"fmt"
	"os"
)

func main() {
	//penerapan fungsi os.Exit()
	defer fmt.Println("hello kalian")
	os.Exit(1)
	fmt.Println("selamat datang di os exit")
}
