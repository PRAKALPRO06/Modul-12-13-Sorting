package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan selection sort
func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		// Tukar elemen
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&n)

	// Loop untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		// Input nomor rumah
		houses := make([]int, m)
		fmt.Printf("Masukkan nomor rumah kerabat di daerah ke-%d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		// Urutkan menggunakan selection sort
		selectionSortDesc(houses)

		// Cetak hasil
		fmt.Printf("Urutan nomor rumah di daerah ke-%d: ", i+1)
		for _, house := range houses {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}
}
