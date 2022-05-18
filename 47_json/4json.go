package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Nama"`
	Age      int
}

func main() {
	var objek = []User{{"saepul bahri", 19}, {"bahri semet", 23}}
	var dataJson, err = json.Marshal(objek)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var jsonstring = string(dataJson)
	fmt.Println(jsonstring)
}
