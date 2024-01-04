// Akmal Aufa Alim
// NRP: 1152200033
// Soal No 4

package main

import (
	"fmt"
)

func main() {
	fmt.Println(polindromAngka(121))
	fmt.Println(polindromAngka(12321))
	fmt.Println(polindromAngka(12345678987654321))
	fmt.Println(polindromAngka(319384913))
	fmt.Println(polindromAngka(4249845))
}

func polindromAngka(n int) bool {
	angkaAsli := n      
	// variabel angka asli yang nanti nya akan di bandingan dengan angka reverse

	var reverseAngka int

	for n > 0 {
		digit := n % 10
		reverseAngka = reverseAngka*10 + digit
		n = n / 10
	}
	// Dalam setiap iterasi, fungsi akan mengambil digit terakhir dari angka yang diberikan, mengalikan dengan 10, dan menambahkan hasilnya ke variabel reverseAngka. Selanjutnya, angka yang diberikan akan dipotongkan oleh 10 agar digit terakhir tidak lagi ada.

	return angkaAsli == reverseAngka
	// Setelah variabel reverseAngka sudah di proses, maka membandingkannya dengan angkaAsli. Jika sama maka True, jika tidak sama maka False.
}
