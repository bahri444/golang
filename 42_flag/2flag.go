package main

import (
	"flag"
	"fmt"
)

func main() {
	var nama = flag.String("nama", "saepul bahri", "type your name")
	var berat = flag.Int64("berat", 65, "type your semester")
	flag.Parse()
	fmt.Printf("nama \t\t:%s\n", *nama)
	fmt.Printf("berat badan \t:%d\n", *berat)
}
