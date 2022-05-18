package main

import (
	"fmt"
	"strings"
)

func main() {
	/*untuk deteksi apakah string yang di cari ada dalam (semua parameter string yang ada). Nilai kembaliannya berupa bool t/f*/
	var data1 = strings.Contains("saepul bahri", "bahri")
	fmt.Println(data1)

	//strings.HasPrefix deteksi apakah sebuah string (parameter pertama) diawali string tertentu (parameter kedua).
	var data2 = strings.HasPrefix("bahri saputra", "bah")
	fmt.Println(data2)

	//deteksi apakah sebuah string (di akhir parameter pertama) diakhiri string tertentu (parameter kedua).
	var data3 = strings.HasSuffix("saepul bahrie", "ie")
	fmt.Println(data3)

	/*menghitung jumlah karakter yang sama sesuai pencarian (parameter kedua)
	dari sebuah string (parameter pertama). Nilai kembalian fungsi ini berupa jumlah
	karakternya*/
	var howMany = strings.Count("bahry saepul", "a")
	fmt.Println(howMany)

	/*mencari posisi indeks sebuah string (parameter kedua) dalam
	string (parameter pertama).Jika diketemukan dua substring, maka yang diambil adalah yang
	pertama,String "n" berada pada indeks 4 dan 8 . Yang dikembalikan adalah yang
	paling kiri (paling kecil), yaitu 4
	*/
	var indeks = strings.Index("bahry saepul", "p")
	fmt.Println(indeks)

	/* mengganti bagian dari string dengan
	string tertentu. Jumlah substring yang di-replace bisa ditentukan, apakah hanya 1
	string pertama, 2 string, atau kesemuanya.*/
	nama := "mangga"
	search := "a"
	ganti := "o"
	deklarasi := strings.Replace(nama, search, ganti, 2)
	fmt.Println(deklarasi)

	//mengulang string (parameter pertama) sebanyak data yang ditentukan (parameter kedua).
	i := strings.Repeat("bahry", 10)
	fmt.Println(i)

	//memisah string (parameter pertama) dengan tanda pemisah bisa ditentukan sendiri (parameter kedua). Hasilnya berupa slice string
	way := strings.Split("bahry semet", ",,,, ")
	fmt.Println(way)

	//string.join
	buah := []string{"jeruk", "mangga", "jeruti"}
	str := strings.Join(buah, "_")
	fmt.Println(str)

	//string.toLower mengubah huruf besar menjadi huruf kecil
	var me = strings.ToLower("BAHRI")
	fmt.Println(me)

	//touper mengubah huruf kecil menjadi huruf besar
	var big = strings.ToUpper("emet aja deh")
	fmt.Println(big)
}
