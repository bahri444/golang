package main

//Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.

import (
	"fmt"
	"math"
)

//deklarsi interface
type hitung interface {
	luas() float64
	keliling() float64
}
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jari() float64 {
	return l.diameter / 2
}
func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jari(), 2)
}
func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}
func (p persegi) keliling() float64 {
	return p.sisi * 4
}
func main() {
	var bangun_datar hitung
	bangun_datar = persegi{10.0}
	fmt.Println("luas :", bangun_datar.luas())
	fmt.Println("keliling :", bangun_datar.keliling())

	bangun_datar = lingkaran{14.5}
	fmt.Println("luas :", bangun_datar.luas())
	fmt.Println("keliling :", bangun_datar.keliling())
	fmt.Println("jari jari :", bangun_datar.(lingkaran).jari())
}
