package main

import (
	"fmt"
	"time"
)

func main() {
	n := 5
	duration := time.Duration(n) * time.Second
	fmt.Println("detik :", duration.Seconds())
	fmt.Println("menit :", duration.Minutes())
	fmt.Println("jam :", duration.Hours())
}
