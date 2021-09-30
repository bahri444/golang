package main

import (
	"fmt"
	"runtime"
)

func getapper(bilangan []int, ch chan float64) {
	var sum = 0
	for _, i := range bilangan {
		sum += i
	}
	ch <- float64(sum) / float64(len(bilangan))
}
func getnum(bilangan []int, ch chan int) {
	var max = bilangan[0]
	for _, i := range bilangan {
		if max < i {
			max = i
		}
	}
	ch <- max
}
func main() {
	runtime.GOMAXPROCS(2)
	var bilangan = []int{1, 2, 3, 2, 3, 6, 7, 8, 9, 6, 5}
	fmt.Println(bilangan)

	var ch1 = make(chan float64)
	go getapper(bilangan, ch1)
	var ch2 = make(chan int)
	go getnum(bilangan, ch2)

	for j := 0; j < 2; j++ {
		select { //Select ini mempermudah kontrol komunikasi data lewat satu ataupun banyak channel.
		case apper := <-ch1: //kondisi apper dan max akan di penuhi jika ada pengiriman data dari ch1 dan ch2
			fmt.Printf("apper :%.2f\n", apper)
		case max := <-ch2:
			fmt.Printf("maxim :%d\n", max)
		}
	}
}
