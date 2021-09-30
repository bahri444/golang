package main

import (
	"fmt"
	"reflect"
)

//Pengaksesan Informasi Method Variabel Objek
type mhs struct {
	Nama  string
	Nilai int
}

func (m *mhs) SetNama(nama string) {
	m.Nama = nama
}
func main() {
	var m1 = &mhs{Nama: "saepul bahri", Nilai: 99}
	fmt.Println("nama anda :", m1.Nama)

	var refVal = reflect.ValueOf(m1)
	var method = refVal.MethodByName("SetNama")
	method.Call([]reflect.Value{
		reflect.ValueOf("bahri"),
	})
	fmt.Println("nama anda :", m1.Nama)
}
