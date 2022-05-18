package main

import (
	"fmt"
	"runtime"
)

//channel sebagai tipe data parameter
func cetakpesan(what chan string) {
	fmt.Println(<-what)
}
func main() {
	runtime.GOMAXPROCS(2)
	var pesan = make(chan string)
	//Iterasi Data Slice/Array Langsung Pada Saat Inisialisasi
	for _, each := range []string{"bahri", "semet", "saepul", "muri"} {
		go func(who string) {
			var data = fmt.Sprintf("hello :%s", who)
			pesan <- data
		}(each)
	}
	for i := 0; i <= 3; i++ {
		cetakpesan(pesan)
	}

}
