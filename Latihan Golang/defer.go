package main

import ("fmt")

func displaysatu() {
	fmt.Println("Satu")
}

func main() {
	defer displaysatu()
	
	// error case
	// num := 0
	// hasil := 1 / num
	// fmt.Println(hasil)
	
	fmt.Println("1")
	fmt.Println("2")
}