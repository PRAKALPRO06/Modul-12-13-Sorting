package main

import (
	"fmt"
)

// Fungsi untuk menghitung median
func calculateMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	} else {
		return (arr[n/2-1] + arr[n/2]) / 2
	}
}

// Fungsi untuk menyisipkan elemen ke dalam array menggunakan insertion sort
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	var num int
	data := []int{}

	for {
		fmt.Scan(&num)

		if num == 0 {
			// Ketika menemukan 0, urutkan dan hitung median
			data = insertionSort(data)
			median := calculateMedian(data)
			fmt.Println(median)
		} else if num == -5313 {
			// Jika menemukan marker akhir, keluar dari loop
			break
		} else {
			// Tambahkan bilangan ke array
			data = append(data, num)
		}
	}
}
