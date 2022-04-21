package main

import "math"

type Segitiga struct {
	Sisi float64
}

func (k Segitiga) Vol() float64 {
	return math.Pow(k.Sisi, 3)
}
func (k Segitiga) Lu() float64 {
	return math.Pow(k.Sisi, 2) * 7
}
func (k Segitiga) Kel() float64 {
	return k.Sisi * 14
}
