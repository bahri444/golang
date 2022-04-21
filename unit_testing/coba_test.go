package main

import "testing"

var (
	segitiga           Segitiga = Segitiga{4}
	volumeSeharusnya   float64  = 64
	luasSeharusnya     float64  = 100
	kelilingSeharusnya float64  = 32
)

func TestVolume(t *testing.T) {
	t.Logf("Volume Kubus : %.2f", segitiga.Vol())
	if segitiga.Vol() != volumeSeharusnya {
		t.Errorf("salah volume seharusnya adalah %.2f", volumeSeharusnya)
	}
}
func TestLuas(t *testing.T) {
	t.Logf("Luas Kubus : %.2f", segitiga.Lu())
	if segitiga.Lu() != luasSeharusnya {
		t.Errorf("salah seharusnya : %.2f", luasSeharusnya)
	}
}
func TestKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", segitiga.Kel())
	if segitiga.Kel() != kelilingSeharusnya {
		t.Errorf("salah keliling seharusnya : %.2f", kelilingSeharusnya)
	}
}
