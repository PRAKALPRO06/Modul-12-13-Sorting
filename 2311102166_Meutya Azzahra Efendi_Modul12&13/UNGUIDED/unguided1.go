// Meutya Azzahra Efendi
// 2311102166
// IF-11-06

package main

import "fmt"

// Fungsi untuk mengurutkan array dengan insertion sort
func insertionSort(arr []int, asc bool) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		// Untuk pengurutan menaik
		if asc {
			for j >= 0 && arr[j] > key {
				arr[j+1] = arr[j]
				j--
			}
		} else { // Untuk pengurutan menurun
			for j >= 0 && arr[j] < key {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
	}
}

// Fungsi untuk menampilkan nomor rumah berdasarkan daerah
func processDaerah(i, m int) ([]int, []int) {
	arr := make([]int, m)
	// Input nomor rumah kerabat untuk setiap daerah
	fmt.Printf("Masukkan %d nomor rumah kerabat untuk daerah ke-%d: ", m, i+1)
	for j := 0; j < m; j++ {
		fmt.Scan(&arr[j])
	}

	// Memisahkan nomor ganjil dan genap
	var odd, even []int
	for _, num := range arr {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return odd, even
}

// Fungsi untuk menampilkan hasil terurut
func printSortedResults(i int, odd, even []int) {
	// Mengurutkan nomor rumah
	insertionSort(odd, true)   // urutkan ganjil secara menaik
	insertionSort(even, false) // urutkan genap secara menurun

	// Output nomor rumah yang terurut
	fmt.Printf("\nNomor rumah terurut untuk daerah ke-%d:\n", i+1)

	// Tampilkan nomor ganjil
	for _, num := range odd {
		fmt.Printf("%d ", num)
	}

	// Tampilkan nomor genap
	for _, num := range even {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// Fungsi utama program
func main() {
	var n int
	// Input jumlah daerah kerabat
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	// Proses untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		// Input jumlah rumah kerabat untuk setiap daerah
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah ke-%d (m): ", i+1)
		fmt.Scan(&m)

		// Ambil nomor rumah untuk daerah tersebut
		odd, even := processDaerah(i, m)

		// Menampilkan hasil terurut untuk daerah tersebut
		printSortedResults(i, odd, even)
	}
}