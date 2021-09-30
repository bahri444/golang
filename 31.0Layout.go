package main

import "fmt"

type mahasiswa struct {
	nama      string
	tinggi    float64
	peringkat int32
	keaadaan  bool
	hobi      []string
}

func main() {
	var data = mahasiswa{
		nama:      "bahry",
		tinggi:    182.5,
		peringkat: 10,
		keaadaan:  false,
		hobi:      []string{"tidur", "belajar ngoding"},
	}
	fmt.Printf("mod b :%b\n", data.peringkat) //'%b'konversi bilangan desimal ke biner

	fmt.Printf("mod c :%c\n", 1400) //'%c'konversi numerik yang merupakan kode unicode,
	fmt.Printf("mod c :%c\n", 1235) //menjadi bentuk string karakter unicode-nya.

	fmt.Printf("mod d :%d\n", data.peringkat) //'%e'konversi numerik menjadi basis bilangan yang kita gunakan

	fmt.Printf("mod e :%e\n", data.tinggi) //'%e'konversi data numerik desimal ke dalam bentuk notasi numerik standar
	fmt.Printf("mod E :%E\n", data.tinggi) //'%E'konversi data numerik desimal ke dalam bentuk notasi numerik standar

	fmt.Printf("mod f :%f\n", data.tinggi)   //'%f'konversi data numerik desimal, dengan lebar desimal bisa ditentukan.
	fmt.Printf("mod f :%.9f\n", data.tinggi) //Secara default lebar digit desimal adalah 6 digit.
	fmt.Printf("mod f :%.2f\n", data.tinggi) //'%f'
	fmt.Printf("mod f :%.f\n", data.tinggi)  //'%f'

	fmt.Printf("mod g :%g\n", 0.123123)       //'%g'konversi lebar digit desimal adalah sesuai dengan datanya aslinya
	fmt.Printf("mod o :%o\n", data.peringkat) //'%o'konversi bentuk string menjadi numerik(oktal)
	fmt.Printf("mod p :%p\n", &data.nama)     //'%p'konversi  data pointer, mengembalikan alamat pointer referensi variabel-nya.
	fmt.Printf("mod q:%q\n", "nama\tinggi")   //'%q'untuk escape string. Meskipun string yang dipakai menggunakan literal \ akan tetap di-escape.
	fmt.Printf("mod s :%s\n", data.nama)      //'%s'konversi string

	fmt.Printf("mod t :%t\n", data.keaadaan) //'%t'konversi data boolean, menampilkan nilai bool-nya
	fmt.Printf("mod T :%T\n", data.nama)     //'%T'mengambil tipe variabel yang akan di konversi

	fmt.Printf("mod v :%v\n", data)   //'%v'hanya mengambil dan menampilkan nilai dari variabel
	fmt.Printf("mod +v :%+v\n", data) //'%+v'menampilkan nama variabel dan nilai variabelnya
	fmt.Printf("mod #v :%#v\n", data) //'%#v'menampilkan nama struct-nya, nama variabel dan nilai variabelnya

	var d = data.nama
	fmt.Printf("mod xxxx :%x%x%x\n", d[0], d[1], d[2]) //'%xxxxx'Jika digunakan pada tipe data string, maka akan mengembalikan kode	heksadesimal tiap karakter
	fmt.Printf("mod x :%x\n", data.peringkat)          //'%x'konversi data numerik menjadi heksa desimal "1a"
	fmt.Printf("mod x :%x\n", d)
	fmt.Printf("mod-mod :%%\n") //'%%'Cara untuk menulis karakter % pada string format.
}
