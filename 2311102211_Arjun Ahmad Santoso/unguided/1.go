package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan Selection Sort
func selectionSortKunjungan(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxEx := i
		minGanjil := -1
		for j := i + 1; j < n; j++ {
			// Cari elemen terkecil
			if (arr[idxEx] % 2) == 1 {
				if (arr[j] % 2) == 1 {
					if arr[j] < arr[idxEx] {
						idxEx = j
					}
				}
			} else {
				if (arr[j] % 2) == 1 {
					if (minGanjil == -1 || arr[j] < minGanjil) {
						idxEx = j
						minGanjil = arr[j]
					}
				} else {
					if arr[j] > arr[idxEx] {
						idxEx = j
					}
				}
			}
		}
		// Tukar elemen terkecil dengan elemen di posisi i
		arr[i], arr[idxEx] = arr[idxEx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	// Proses tiap daerah
	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		// Membaca nomor rumah untuk daerah ini
		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		// Urutkan array dari terkecil ke terbesar
		selectionSortKunjungan(arr, m)

		// Tampilkan hasil
		fmt.Printf("Nomor rumah terurut untuk dikunjungi %d: ", daerah)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
