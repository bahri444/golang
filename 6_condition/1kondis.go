//contoh kondisi statement
package main

import "fmt"

func main() {
	var bil = 9

	if bil == 10 {
		fmt.Println("nilai A")
	} else if bil >= 8 {
		fmt.Println("nilai B")
	} else if bil >= 7 {
		fmt.Println("nilai c")
	} else {
		fmt.Println("tidak lulus")
	}
}
