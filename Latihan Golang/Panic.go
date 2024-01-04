package main

import ("fmt")

func pembagian (angka1, angka2 int) {
	if angka2 == 0 {
		panic("Jangan panik, tidak bisa pembagian dengan nol")
	} else {
		hasil := angka1 / angka2
		fmt.Println("Hasil nya =", hasil)
	}
}

func main(){
	var angka1, angka2 int
	fmt.Print("Masukan Angka 1: ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukan Angka 2: ")
	fmt.Scanln(&angka2)

	pembagian(angka1, angka2)
}