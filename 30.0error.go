package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("input data Number")
	fmt.Scanln(&input) //untuk men-capture inputan yang diketik oleh user sebelum dia menekan enter,
	//lalu menyimpannya sebagai string ke variabel input

	var bil int
	var err error                  //error adalah tipe data dari variabel err
	bil, err = strconv.Atoi(input) //mengembalikan data yang ti tampung oleh bil dan err
	if err == nil {                //ketika konversi berjalan mulus maka nilai balik kedua adalah nil.
		//jika gagal penyebabnya bisa langsung di ketahui

		fmt.Println(bil, "done integer")
	} else {
		fmt.Println(input, "not integer")
		fmt.Println(err.Error()) //error memiliki 1 buah properti berupa method "Error()"
	}
}
