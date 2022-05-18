package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Nama"`
	Alamat   string
	Nilai    int
}

func main() {
	var stringjson = `[{"Nama":"Saepul Bahri","Alamat":"melar","Nilai":90},{"Nama":"Bahri semet","Alamat":"selebung","Nilai":88}]`
	var data []User
	err := json.Unmarshal([]byte(stringjson), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Nama :", data[0].FullName)
	fmt.Println("Alamat :", data[0].Alamat)
	fmt.Println("Nilai :", data[0].Nilai)
	fmt.Println()
	fmt.Println("Nama :", data[1].FullName)
	fmt.Println("Alamat :", data[1].Alamat)
	fmt.Println("Nilai :", data[1].Nilai)

}
