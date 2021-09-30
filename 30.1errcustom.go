package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("tidak boleh kosong")
	}
	return true, nil
}
func main() {
	var nama string
	fmt.Println("silahkan masukkan nama anda :")
	fmt.Scanln(&nama)
	if valid, err := validate(nama); valid {
		fmt.Println("nama anda adalah :", nama)
	} else {
		fmt.Println(err.Error())
	}
}
