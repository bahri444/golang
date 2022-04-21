package main

import "testing"

var (
	kubus        Kubus   = Kubus{4}
	volumeAsli   float64 = 64
	luasAsli     float64 = 96
	kelilingAsli float64 = 48
)

func TestHitungVol(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Volume())
	if kubus.Volume() != volumeAsli {
		t.Errorf("Salah seharusnya adalah : %.2f", volumeAsli)
	}
}
func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", kubus.Luas())
	if kubus.Luas() != luasAsli {
		t.Errorf("Salah seharusnya adalah : %.2f", luasAsli)
	}
}
func TestHitungKel(t *testing.T) {
	t.Logf("Keliling : %.2f", kubus.Keliling())
	if kubus.Keliling() != kelilingAsli {
		t.Errorf("Salah seharusnya adalah : %.2f", kelilingAsli)
	}
}

// func main() {

// }
