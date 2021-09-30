package main

import (
	"fmt"
	"os"
	"time"
)

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}
func wetcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out no answer and more than ", timeout, "second")
	os.Exit(0)
}
func main() {
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go wetcher(timeout, ch)

	var input string
	fmt.Print("input data :")
	fmt.Scanln(&input)
	if input == "29" {
		fmt.Println("the answer is right")

	} else {
		fmt.Println("the answer is wrong")
	}
}
