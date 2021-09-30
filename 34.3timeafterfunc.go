package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan bool)
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("berhenti sejenak")
		ch <- true
	})
	fmt.Println("mulai eksekusi")
	<-ch
	fmt.Println("eksekusi selesai")

	//time after
	fmt.Println("\nmemulai time.after")
	<-time.After(4 * time.Second)
	fmt.Println("berhentilah")
	fmt.Println("selesai")
}
