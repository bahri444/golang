package main

import (
	"fmt"
	"time"
)

func main() {
	var data, _ = time.Parse(time.RFC822, "02 Sep 21 08:00 WIB") //value
	var datas1 = data.Format("Monday 02 January 2006 15:04 MST") //format
	fmt.Println("datas1=", datas1)

	var datas2 = data.Format(time.RFC3339)
	fmt.Println("datas2=", datas2)
}
