package main

import "fmt"

type mhs struct {
	nama string
}

func (m mhs) changenama1(nama string) {
	fmt.Println("change nama 1 to :", nama)
	m.nama = nama
}
func (m *mhs) changenama2(nama string) {
	fmt.Println("change nama 2 to :", nama)
	m.nama = nama
}
func main() {
	var m1 = mhs{"saepul bahri"}
	fmt.Println("nama1 before :", m1.nama)

	m1.changenama1("bang emet")
	fmt.Println("m1 nama 1 after change :", m1.nama) /*Setelah eksekusi statement s1.changeName1("bang emet") , nilai s1.name tidak
	berubah. Sebenarnya nilainya berubah tapi hanya dalam method changeName1()
	saja, nilai pada reference di objek-nya tidak berubah. Karena itulah ketika objek di
	print value dari s1.name tidak berubah.*/

	m1.changenama2("saputra bahri")
	fmt.Println("m1 nama 2 after change :", m1.nama)

}
