package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fulldate string `json:"Nama"`
	Alamat   string
	Nim      int
}

func main() {
	var jsondate = `{"Nama":"saepul bahri","Alamat":"Melar desa selebung","Nim":17200048}`
	var jsoncasting = []byte(jsondate)
	var Data User
	var err = json.Unmarshal(jsoncasting, &Data)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Nama :", Data.Fulldate)
	fmt.Println("Alamat :", Data.Alamat)
	fmt.Println("Nim :", Data.Nim)
}
