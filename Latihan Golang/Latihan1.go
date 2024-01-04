// package main

// import (
// 	"fmt"
// 	// "strconv"
// )

// // func main() {
// // 	var fahrenheit float32
// // 	fmt.Print("Masukan Fahrenheit: ")
// // 	fmt.Scanln(&fahrenheit)

// // 	fmt.Println("Suhu F: ",fahrenheit)
// // 	var c = (fahrenheit - 32) * 5/9
// // 	fmt.Println("Suhu C:",c)
// // }

// // package main

// // func main() {
// // 	var ft float64
// // 	fmt.Print("Masukan berapa feet: ")
// // 	fmt.Scanln(&ft)

// // 	fmt.Println("Feet:", ft)
// // 	var hasil = ft * 0.3048
// // 	fmt.Println(ft, "ft:", hasil)
// // }

// // package main

// // import (
// // 	"fmt"
// // 	"strconv"
// // )

// // func main() {
// // 	var number int

// // 	// Minta pengguna memasukkan angka
// // 	fmt.Print("Masukkan angka: ")
// // 	fmt.Scanln(&number)

// // 	if isPalindrome(number) {
// // 		fmt.Printf("%d adalah palindrome.\n", number)
// // 	} else {
// // 		fmt.Printf("%d bukan palindrome.\n", number)
// // 	}
// // }

// // Fungsi untuk memeriksa apakah sebuah angka adalah palindrome
// // func isPalindrome(number int) bool {
// // 	// Konversi angka ke string
// // 	numStr := strconv.Itoa(number)

// // 	// Panjang string
// // 	length := len(numStr)

// // 	// Periksa karakter pertama dengan karakter terakhir, kedua dengan kedua dari terakhir, dan seterusnya
// // 	for i := 0; i < length/2; i++ {
// // 		if numStr[i] != numStr[length-1-i] {
// // 			return false
// // 		}
// // 	}
// // 	return true
// // }


// func removeElement(nums []int, val int) int {
//     // Inisialisasi indeks untuk menyimpan elemen yang tidak sama dengan val
//     idx := 0

//     // Iterasi melalui array nums
//     for _, num := range nums {
//         if num != val {
//             // Jika elemen tidak sama dengan val, maka tambahkan ke array
//             nums[idx] = num
//             idx++
//         }
//     }

//     return idx
// }

// func main() {
//     // Contoh penggunaan
//     nums1 := []int{3, 2, 2, 3}
//     val1 := 2
//     result1 := removeElement(nums1, val1)
//     fmt.Println("Output 1:", result1, "nums 1:", nums1[:result1])

//     nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
//     val2 := 4
//     result2 := removeElement(nums2, val2)
//     fmt.Println("Output 2:", result2, "nums 2:", nums2[:result2])
// }

// func mySqrt(x int) int {
//     if x == 0 {
//         return 0
//     }

//     // Inisialisasi batasan bawah dan batasan atas untuk pencarian biner
//     left, right := 1, x

//     for left <= right {
//         mid := left + (right-left)/2 // Menghindari integer overflow
//         square := mid * mid

//         if square == x {
//             return mid
//         } else if square < x {
//             left = mid + 1
//         } else {
//             right = mid - 1
//         }
//     }

//     return right
// }

// func main() {
//     // Contoh penggunaan
//     x1 := 21
//     result1 := mySqrt(x1)
//     fmt.Println("Output 1:", result1)

//     x2 := 8
//     result2 := mySqrt(x2)
//     fmt.Println("Output 2:", result2)
// }

