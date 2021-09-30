package main

import "fmt"

func main() {
	type mhs struct {
		nama  string
		nilai int
	}
	var data interface{} = &mhs{"saepul bahri", 80}
	var nama = data.(*mhs).nama //casting variabel interface kosong ke objek pointer
	var nilai = data.(*mhs).nilai
	fmt.Println("nama anda :", nama)
	fmt.Println("nilai anda:", nilai)

}
