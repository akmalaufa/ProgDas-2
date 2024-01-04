// Kelompok Botak Keren
// Nama Anggota kelompok:
// 1. Akmal Aufa Alim 1152200033
// 2. Dioba Rizky 1152200008
// 3. Satrio Andharbenny 1152200012

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

var bandaraData = map[string]string{
	"CGK": "Bandar Udara Halim Perdanakusuma, Jakarta",
	"SUB": "Bandar Udara Internasional Supadio, Pontianak",
	"DPS": "Bandar Udara Internasional Sultan Mahmud Badaruddin II, Palembang",
	"KNO": "Bandar Udara Internasional Minangkabau, Padang",
}

var tarifData = map[string]int{
	"CGK_SUB": 1000000,
	// "SUB_CGK": 1000000,
	"CGK_DPS": 1500000,
	// "DPS_CGK": 1500000,
	"CGK_KNO": 1200000,
	// "KNO_CGK": 1200000,
	"SUB_DPS": 800000,
	// "DPS_SUB": 800000,
	"KNO_SUB": 900000,
	// "SUB_KNO": 900000,
	"KNO_DPS": 700000,
	// "DPS_KNO": 700000,
}

func hitungBiaya(asal, tujuan string, jumlahPenumpang int, jenisTiket string) (int, int) {
	tarif, ok := tarifData[asal+"_"+tujuan]
	if !ok {
		return 0, 0
	}

	switch jenisTiket {
	case "ekonomi":
		tarif *= 1
	case "bisnis":
		tarif *= 2
	case "firstclass":
		tarif *= 3
	}

	totalBiaya := tarif * jumlahPenumpang
	kodeUnik := rand.Intn(900) + 100

	return totalBiaya, kodeUnik
}

func prosesPembayaran(totalBiaya, kodeUnik int, metodePembayaran string) int {
	switch metodePembayaran {
	case "transfer":
		return totalBiaya + kodeUnik
	case "kredit":
		diskon := 0.02
		totalPembayaran := int(float64(totalBiaya)*(1-diskon)) + kodeUnik
		return totalPembayaran
	default:
		fmt.Println("Metode pembayaran tidak valid.")
		return 0
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var namaPemesan, asalBandara, tujuanBandara, tanggalKeberangkatan, tanggalPulang, pulangPergi, jenisTiket, metodePembayaran string
	var jumlahPenumpang int

	fmt.Print("Masukkan Nama Pemesan: ")
	fmt.Scanln(&namaPemesan)
	fmt.Println(" ")
	fmt.Println("Kode Bandara")
	fmt.Println("CGK: Bandar Udara Halim Perdanakusuma, Jakarta")
	fmt.Println("SUB: Bandar Udara Internasional Supadio, Pontianak")
	fmt.Println("DPS: Bandar Udara Internasional Sultan Mahmud Badaruddin II, Palembang")
	fmt.Println("KNO: Bandar Udara Internasional Minangkabau, Padang")
	fmt.Print("Masukkan Asal Bandara (Kode): ")
	fmt.Scanln(&asalBandara)
	fmt.Print("Masukkan Tujuan Bandara (Kode): ")
	fmt.Scanln(&tujuanBandara)
	fmt.Print("Masukkan Tanggal Keberangkatan (YYYY-MM-DD): ")
	fmt.Scanln(&tanggalKeberangkatan)
	fmt.Print("Pilih Tipe Penerbangan (pulang/pergi): ")
	fmt.Scanln(&pulangPergi)

	if pulangPergi == "pulang" {
		fmt.Print("Masukkan Tanggal Kepulangan (YYYY-MM-DD): ")
		fmt.Scanln(&tanggalPulang)
	}

	fmt.Print("Masukkan Jumlah Penumpang: ")
	fmt.Scanln(&jumlahPenumpang)
	fmt.Print("Masukkan Jenis Tiket (ekonomi/bisnis/firstclass): ")
	fmt.Scanln(&jenisTiket)
	fmt.Print("Pilih Metode Pembayaran (transfer/kredit): ")
	fmt.Scanln(&metodePembayaran)

	var totalBiaya, kodeUnik int
	if pulangPergi == "pulang" {
		totalBiaya, kodeUnik = hitungBiaya(asalBandara, tujuanBandara, jumlahPenumpang, jenisTiket)
	} else {
		totalBiayaKeberangkatan, _ := hitungBiaya(asalBandara, tujuanBandara, jumlahPenumpang, jenisTiket)
		totalBiayaPulang, _ := hitungBiaya(tujuanBandara, asalBandara, jumlahPenumpang, jenisTiket)
		totalBiaya = totalBiayaKeberangkatan + totalBiayaPulang
		kodeUnik = rand.Intn(900) + 100
	}

	totalPembayaran := prosesPembayaran(totalBiaya, kodeUnik, metodePembayaran)

	fmt.Println("\nInformasi Pemesanan:")
	fmt.Printf("Nama Pemesan: %s\n", namaPemesan)
	fmt.Printf("Asal Bandara: %s\n", bandaraData[asalBandara])
	fmt.Printf("Tujuan Bandara: %s\n", bandaraData[tujuanBandara])
	fmt.Printf("Tanggal Keberangkatan: %s\n", tanggalKeberangkatan)
	if pulangPergi == "pulang" {
		fmt.Printf("Tanggal Kepulangan: %s\n", tanggalPulang)
	}
	fmt.Printf("Jumlah Penumpang: %d\n", jumlahPenumpang)
	fmt.Printf("Jenis Tiket: %s\n", jenisTiket)
	fmt.Printf("Total Biaya: Rp. %d + %d\n", totalBiaya, kodeUnik)
	fmt.Printf("Metode Pembayaran: %s\n", metodePembayaran)
	fmt.Printf("Total Pembayaran: Rp. %d\n", totalPembayaran)
	terbilang := Terbilang(totalPembayaran)
	fmt.Printf("Rp. %d ---> %s Rupiah\n", totalPembayaran, strings.Title(terbilang)) // menggunakan strings.Title untuk mengubah huruf pertama menjadi huruf besar

}
