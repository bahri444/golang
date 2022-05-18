package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}
func catch() {
	if r := recover(); r != nil {
		fmt.Println("error occured", r)
	} else {
		fmt.Println("application running perfect")
	}
}
func main() {
	defer catch()

	var nama string
	fmt.Print("silahkan input data :")
	fmt.Scanln(&nama)
	if valid, err := validate(nama); valid {
		fmt.Println("hello", nama)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
