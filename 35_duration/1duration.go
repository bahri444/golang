package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	duration := time.Since(start)
	fmt.Println("time Second :", duration.Seconds())
	fmt.Println("time Minutes :", duration.Minutes())
	fmt.Println("time Hours :", duration.Hours())
}
