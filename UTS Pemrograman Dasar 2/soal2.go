// Nama: Akmal Aufa Alim
// NRP : 1152200033
// Soal no 2

package main

import (
	"fmt"
	"strings"
)

func hitungKonsonanVokal(kalimat string) (int, int) {
	konsonan := 0
	vokal := 0
	// inisilasasi variabel konsonan dan lokal dengan 0

	for _, c := range strings.ToLower(kalimat) {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vokal++
		case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
			konsonan++
		}
	}
	// Iterasikan kalimat 1 huruf per huruf lalu bandingkan setiap huruf nya dengan case lalu jika vokal/konsonan tambahkan ke variabel nya
	
	return konsonan, vokal
	// Balikan nilai hasil jumlah huruf konsonan dan vokal
}

func main() {
	kalimat := "Akmal Aufa Alim"
	
	konsonan, vokal := hitungKonsonanVokal(kalimat)
	
	fmt.Println("Jumlah huruf vokal =", vokal)
	fmt.Println("Jumlah huruf konsonan =", konsonan)
}