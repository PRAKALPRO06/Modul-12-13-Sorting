package main

import (
	"fmt"
)

// Erwin Rivaldo Silaban 2311102248
func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		// untuk menukar elemen terkecil dengan elemen di posisi i
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func median(arr []int, n int) int {
	if n%2 == 0 {
		return (arr[n/2-1] + arr[n/2]) / 2
	} else {
		return arr[(n-1)/2]
	}
}

func main() {
	var arr [1000000]int
	var arrMedian [1000000]int
	var n, nMed, datum int

	fmt.Printf("Masukkan data: ")
	for datum != -5313 && n < 1000000 {
		fmt.Scan(&datum)
		if datum == 0 {
			selectionSort(arr[:n], n)
			arrMedian[nMed] = median(arr[:n], n)
			nMed++
		} else {
			arr[n] = datum
			n++
		}
	}

	fmt.Println(" Hasil: ")
	for i := 0; i < nMed; i++ {
		fmt.Println(arrMedian[i], " ")
	}
}
