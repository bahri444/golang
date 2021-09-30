package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var url = "https://bahrysaipulsemet.com"
	var encode = base64.URLEncoding.EncodeToString([]byte(url))
	var encodestr = string(encode)
	fmt.Println(encodestr)
	var decodebyte, _ = base64.URLEncoding.DecodeString(encode)
	var decodestr = string(decodebyte)
	fmt.Println(decodestr)
}
