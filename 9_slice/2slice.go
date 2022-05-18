package main

import "fmt"

func main() {
	var buahan = []string{"jeruk", "kelapa", "anggur", "semangka"}
	//len mengambil element dari index ke nol hingga ke element ke n-1
	//cap mengambil element dari index ke satuhingga ke selulruh isi element
	fmt.Println(len(buahan))
	fmt.Println(cap(buahan)) //cap menghitung seluruh kapasitas/element pada slice

	var abuahan = buahan[0:3]
	fmt.Println(len(abuahan))
	fmt.Println(cap(abuahan))

	var bbuahan = buahan[1:4]
	fmt.Println(len(bbuahan))
	fmt.Println(cap(bbuahan))

}
