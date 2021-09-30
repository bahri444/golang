package main

import "fmt"

func main() {
	var lembu map[string]int
	lembu = map[string]int{}
	lembu["januari"] = 10
	lembu["februari"] = 5
	fmt.Println("jumlah lembu di bulan janunari", lembu["januari"])
	fmt.Println("jumlah lembu di bulan april", lembu["april"]) //hasilnya 0 karena blan april tidak terdaftar

}
