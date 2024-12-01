package main

import "fmt"

// Fungsi untuk mengurutkan array menggunakan Selection Sort
func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			// Cari elemen terkecil
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		// Tukar elemen terkecil dengan elemen di posisi i
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

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

		// Pisahkan nomor ganjil dan genap ke dalam dua slice berbeda
		var ganjil, genap []int
		for _, num := range arr {
			if num%2 == 1 {
				ganjil = append(ganjil, num)
			} else {
				genap = append(genap, num)
			}
		}

		// Urutkan nomor ganjil membesar dan nomor genap mengecil
		selectionSort(ganjil, len(ganjil))
		selectionSort(genap, len(genap))

		// Reverse nomor genap agar urutannya mengecil
		for i, j := 0, len(genap)-1; i < j; i, j = i+1, j-1 {
			genap[i], genap[j] = genap[j], genap[i]
		}

		// Gabungkan dan tampilkan hasil
		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
