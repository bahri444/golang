package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("data input tidak boleh kosong")
	}
	return true, nil
}
func main() {
	var nama string
	fmt.Print("input data :")
	fmt.Scanln(&nama)
	if valid, err := validate(nama); valid {
		fmt.Println("nama anda ", nama)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
