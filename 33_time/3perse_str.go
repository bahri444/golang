package main

import (
	"fmt"
	"time"
)

func main() {
	var layoutFormat, value string
	var data time.Time
	layoutFormat = "2006-01-02 15:04:05" //format default menentukan date time secara manual
	value = "2021-06-16 08:04:00"
	data, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", data.String())

	layoutFormat = "02/01/2006 MST"
	value = "16/07/2021 WIB"
	data, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", data.String())
}
