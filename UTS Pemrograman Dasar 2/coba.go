package main

import (
	"fmt"
	"strings"
)

var satuan = []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh"}
var belasan = []string{"", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas", "enam belas", "tujuh belas", "delapan belas", "sembilan belas"}
var puluhan = []string{"", "sepuluh", "dua puluh", "tiga puluh", "empat puluh", "lima puluh", "enam puluh", "tujuh puluh", "delapan puluh", "sembilan puluh"}

func Terbilang(n int) string {
	if n < 0 || n > 999999999 {
		return "Input diluar rentang (0-999999999)"
	}

	if n == 0 {
		return "nol"
	}

	return terbilangSatuan(n)
}

func terbilangSatuan(n int) string {
	if n <= 10 {
		return satuan[n]
	} else if n < 20 {
		return belasan[n-10]
	} else if n < 100 {
		return puluhan[n/10] + " " + terbilangSatuan(n%10)
	} else if n < 1000 {
		return satuan[n/100] + " ratus " + terbilangSatuan(n%100)
	} else if n < 1000000 {
		if n%1000 == 0 {
			return terbilangSatuan(n/1000) + " ribu"
		} else {
			return terbilangSatuan(n/1000) + " ribu " + terbilangSatuan(n%1000)
		}
	} else {
		if n%1000000 == 0 {
			return terbilangSatuan(n/1000000) + " juta"
		} else {
			return terbilangSatuan(n/1000000) + " juta " + terbilangSatuan(n%1000000)
		}
	}
}

func main() {
	angka := 87
	terbilang := Terbilang(angka)
	fmt.Printf("%d ---> %s rupiah\n", angka, strings.Title(terbilang)) // menggunakan strings.Title untuk mengubah huruf pertama menjadi huruf besar
}
