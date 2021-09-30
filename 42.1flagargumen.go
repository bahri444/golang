package main

import (
	"fmt"
	"os"
)

func main() {
	var argsrow = os.Args
	fmt.Printf("=> %#v\n", argsrow)

	var datarow = argsrow[1:]
	fmt.Printf("=> %#v\n", datarow)
}
