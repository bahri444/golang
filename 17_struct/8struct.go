package main

import "fmt"

type mhs struct {
	nama  string
	nilai int
}

func main() {
	//Inisialisasi Slice Anonymous Struct
	//Anonymous struct bisa dijadikan sebagai tipe sebuah slice. Dan nilai awalnya juga bisa diinisialisasi langsung pada saat deklarasi.
	var allmhs = []struct {
		nim int
		mhs
	}{
		{mhs: mhs{"saepul", 80}, nim: 17200048},
		{mhs: mhs{"bahri", 80}, nim: 17200048},
		{mhs: mhs{"emet", 80}, nim: 17200048},
		{mhs: mhs{"bangepul", 80}, nim: 17200048},
		{mhs: mhs{"bang emet", 80}, nim: 17200048},
	}
	for _, mhs := range allmhs {
		fmt.Println(mhs)
	}

}
