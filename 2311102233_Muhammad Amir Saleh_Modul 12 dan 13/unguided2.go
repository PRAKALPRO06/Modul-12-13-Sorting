package main

import (
	"fmt"
)

func selectionSort(arr []int, descending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		selectedIdx := i
		for j := i + 1; j < n; j++ {
			if (descending && arr[j] > arr[selectedIdx]) || (!descending && arr[j] < arr[selectedIdx]) {
				selectedIdx = j
			}
		}
		arr[i], arr[selectedIdx] = arr[selectedIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Jumlah daerah: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Print("Jumlah kerabat: ")
		fmt.Scan(&m)

		rumah := make([]int, m)
		ganjil := []int{}
		genap := []int{}

		fmt.Print("Masukkan nomor rumah: ")
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
			if rumah[j]%2 != 0 {
				ganjil = append(ganjil, rumah[j])
			} else {
				genap = append(genap, rumah[j])
			}
		}

		selectionSort(ganjil, false)
		selectionSort(genap, true)

		for _, r := range ganjil {
			fmt.Print(r, " ")
		}
		for _, r := range genap {
			fmt.Print(r, " ")
		}
		fmt.Println()
	}
}
