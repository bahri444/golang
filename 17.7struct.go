//kombinasi slice dan struct
package main

import "fmt"

type mhs struct {
	nama  string
	nilai int
}

func main() {
	var allmhs = []mhs{{nama: "saepul", nilai: 99},
		{nama: "emet", nilai: 90},
		{nama: "bahri", nilai: 98},
		{nama: "bang epul", nilai: 79},
	}
	for _, mhs := range allmhs {
		fmt.Println("nama :", mhs.nama, "\tis\t", "nilai ;", mhs.nilai)
	}
}
