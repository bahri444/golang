package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Now()
	time.Sleep(2 * time.Second)
	time2 := time.Now()
	selisih := time2.Sub(time1)
	fmt.Println("selisih detik :", selisih.Seconds())
	fmt.Println("selisih menit :", selisih.Minutes())
	fmt.Println("selisih jam   :", selisih.Hours())
}
