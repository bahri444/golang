package main

import "fmt"

func main() {
	//return untuk menghentikan proses dalam fungsi
	dividenumber(10, 2)
	dividenumber(4, 0)
	dividenumber(-4, 2)
}
func dividenumber(m, n int) {
	if n == 0 {
		fmt.Printf("number valid. %d number tidak valid%d \n", m, n)
		return
	}
	var res = m / n //angka yang ada pada func divdenumber di bagi melalui variabel m bagi n
	fmt.Printf("%d %d=%d\n", m, n, res)
}
