package main

import "fmt"

func main() {
	//penerapan if else
	var bil = 88.8820
	if percent := bil / 100; percent >= 90 {
		fmt.Printf("nila A%.1f%s\n", percent, "%")
	} else if bil >= 80 {
		fmt.Printf("nilai B %.1f%s\n", percent, "%")
	} else if bil >= 70 {
		fmt.Printf("nilai C %.1f%s\n", percent, "%")
	} else {
		fmt.Printf("nilai di bawah standar", percent, "%")
	}
}
