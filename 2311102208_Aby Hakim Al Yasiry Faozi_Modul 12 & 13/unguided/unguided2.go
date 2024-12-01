package main

import "fmt"

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

func median(arr []int) int {
	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	}
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	var data []int
	var input int

	fmt.Println("Masukkan data:")

	for {
		fmt.Scan(&input)
		if input == -531351 {
			break
		}

		if input == 0 {
			if len(data) > 0 {
				temp := make([]int, len(data))
				copy(temp, data)
				selectionSort(temp)
				fmt.Println("Median saat ini:", median(temp))
			} else {
				fmt.Println("Tidak ada data untuk dihitung.")
			}
		} else {
			data = append(data, input)
		}
	}
}
