package main

import "fmt"

func main() {
	//if bersarang dengan switch case
	var bil = 2
	if bil >= 9 {
		switch bil {
		case 10:
			fmt.Println("nilai sangat bagus")
		default:
			fmt.Println("nice")
		}
	} else {
		if bil == 5 {
			fmt.Println("nilai anda lima")
		} else if bil == 4 {
			fmt.Println("nilai anda empat")
		} else if bil == 3 {
			fmt.Println("nilai anda tiga")
		} else {
			fmt.Println("nilai anda sangat kurang")
		}
	}
}
