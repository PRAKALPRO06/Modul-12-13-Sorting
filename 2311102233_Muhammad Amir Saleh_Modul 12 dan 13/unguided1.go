package main

import (
	"fmt"
)

func selectionSort(arr []int) {
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

func main() {
	var n int
	fmt.Print("Jumlah daerah: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Print("Jumlah kerabat: ")
		fmt.Scan(&m)

		rumah := make([]int, m)
		fmt.Print("Masukkan nomor rumah: ")
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah)

		for _, r := range rumah {
			fmt.Print(r, " ")
		}
		fmt.Println()
	}
}
