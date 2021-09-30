package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlSTR = "http://bahri.com:80/hello?Nama=Saepul_Bahri&Age=19"
	var url, err = url.Parse(urlSTR)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("url : %s\n", urlSTR)
	fmt.Printf("protocol :%s\n", url.Scheme)
	fmt.Printf("host :%s\n", url.Host)
	fmt.Printf("folder :%s\n", url.Path)

	var Nama = url.Query()["Nama"][0]
	var Usia = url.Query()["Age"][0]
	fmt.Printf("Nama  :%s\nUsia :%s\n", Nama, Usia)
}
