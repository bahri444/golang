package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var nama = "saepul bahry semet"
	var encode = make([]byte, base64.StdEncoding.EncodedLen(len(nama)))
	base64.StdEncoding.Encode(encode, []byte(nama))
	var encodestr = string(encode)
	fmt.Println(encodestr)
	var decode = make([]byte, base64.StdEncoding.DecodedLen(len(encode)))
	var _, err = base64.StdEncoding.Decode(decode, encode)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodestr = string(decode)
	fmt.Println(decodestr)
}
