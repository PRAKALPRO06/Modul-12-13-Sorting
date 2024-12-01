package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku
type Daerah struct {
	ganjil []int
	genap  []int
}

func selectionSortAscending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func selectionSortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&n)

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		var arr []int
		fmt.Print("Masukkan nomor rumah: ")
		for i := 0; i < m; i++ {
			var nomor int
			fmt.Scan(&nomor)
			arr = append(arr, nomor)
		}

		var daerahData Daerah
		for _, num := range arr {
			if num%2 == 1 {
				daerahData.ganjil = append(daerahData.ganjil, num)
			} else {
				daerahData.genap = append(daerahData.genap, num)
			}
		}

		selectionSortAscending(daerahData.ganjil)

		selectionSortDescending(daerahData.genap)

		// Output hasil
		fmt.Printf("Daerah %d:\n", daerah)
		if daerah != 3 {
			for _, num := range daerahData.ganjil {
				fmt.Print(num, " ")
			}
		}
		for _, num := range daerahData.genap {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}
