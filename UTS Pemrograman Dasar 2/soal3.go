// Nama: Akmal Aufa Alim
// NRP : 1152200033
// Soal 3

package main

import (
	"fmt"
)

func cari_target(arr []int, target int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

func main() {
	list := []int{2, 7, 11, 15}
	target := 17
	indikasi := cari_target(list, target)
	if indikasi != nil {
		fmt.Println("Output:", indikasi)
	} else {
		fmt.Println("Tidak ada pasangan bilangan yang menghasilkan total target yang diinginkan")
	}
}