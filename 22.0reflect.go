package main

import (
	"fmt"
	"reflect"
)

//mencari tipe data dan variabel menggunakan reflect
func main() {
	var nilai = 259
	var refVal = reflect.ValueOf(nilai)
	fmt.Println("type variabel :", refVal.Type())

	if refVal.Kind() == reflect.Int {
		fmt.Println("nilai dari variabel :", refVal.Int()) //atau bisa juga menggunakan fmt.Println("nilai dari variabel :", refVal.Interface().(int))
	}
}
