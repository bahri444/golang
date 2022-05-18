package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Full string `json:"Nama"`
	Age  int
}

func main() {
	var jsonString = `{"Nama":"Saepul bahri","Age":20}`
	var castjson = []byte(jsonString)

	var dataone map[string]interface{}
	json.Unmarshal(castjson, &dataone)

	fmt.Println("nama :", dataone["Nama"])
	fmt.Println("usia :", dataone["Age"])

	var datatwo interface{}
	json.Unmarshal(castjson, &datatwo)

	var decodestr = datatwo.(map[string]interface{})
	fmt.Println("Nama :", decodestr["Nama"])
	fmt.Println("Usia :", decodestr["Age"])
}
