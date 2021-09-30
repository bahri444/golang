package main

import "fmt"

func main() {
	//copy berfungsi untuk mencopy element isi dari slice variabel src
	dek := make([]string, 3) //cara deklarasi untuk mencopy isi data sesuai jumlah yg di tentukan(3)
	data := []string{"saepul", "bahri", "saputra", "emet"}
	n := copy(dek, data) //variabel n berfungsi untuk mencetak jumlah element yg di copy
	fmt.Println(dek)
	fmt.Println(data)
	fmt.Println(n)

	//mengganti isi variabel menggunakan copy
	snack := []string{"kentang", "kentang", "kentang", "kentang", "kentang"}
	manipulasi := []string{"bakso", "mie ayam", "soto", "esteler"}
	m := copy(snack, manipulasi)
	fmt.Println(snack)
	fmt.Println(manipulasi)
	fmt.Println(m)
}
