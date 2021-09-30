package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "saepul bahri"
	var encode = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encode :", encode)
	var decodebyte, _ = base64.StdEncoding.DecodeString(encode)
	var str = string(decodebyte)
	fmt.Println("decode :", str)
}
