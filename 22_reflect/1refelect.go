package main

import (
	"fmt"
	"reflect"
)

type mhs struct {
	Nama  string
	Nilai int
}

func (m *mhs) getProperty() {
	var refValue = reflect.ValueOf(m)
	if refValue.Kind() == reflect.Ptr {
		refValue = refValue.Elem()
	}
	var refType = refValue.Type()
	for i := 0; i < refValue.NumField(); i++ {
		fmt.Println("nama variabel :", refType.Field(i).Name)
		fmt.Println("nama variabel :", refType.Field(i).Type)
		fmt.Println("nama variabel :", refValue.Field(i).Interface())
		fmt.Println()
	}
}
func main() {
	//Pengaksesan Informasi Property Variabel Objek
	var m1 = &mhs{Nama: "bahri saputra", Nilai: 99}
	m1.getProperty()
}
