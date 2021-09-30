package main

import (
	"fmt"
	"time"
)

func main() {
	var data = time.Now()
	fmt.Println("Tahun :", data.Year(), "\nbulan :", data.Month())
	fmt.Println("hari :", data.Weekday(), "\ntanggal :", data.Day())
	fmt.Print("jam : ")
	fmt.Println(data.Clock())
}
