package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func senddata(ch chan<- int) {
	for i := 1; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%i+1) * time.Second)
	}
}
func terimadata(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print("terima data", data, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("tidak ada aktifitas setelah 5 detik")
			break loop
		}
	}
}
func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var pesan = make(chan int)

	go senddata(pesan)
	terimadata(pesan)

}
