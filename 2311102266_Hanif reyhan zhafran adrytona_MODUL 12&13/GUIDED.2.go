// 2311102266_Hanif Reyhan Zhafran Arytona

package main

import (
	"fmt"
)

func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func isConstantDifference(arr []int, n int) (bool, int) {
	if n < 2 {
		return true, 0
	}
	difference := arr[1] - arr[0]
	for i := 1; i < n-1; i++ {
		if arr[i+1]-arr[i] != difference {
			return false, 0
		}
	}
	return true, difference
}

func main() {
	var arr []int
	var num int

	fmt.Println("Masukkan data integer (akhiri dengan bilangan negatif):")
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		arr = append(arr, num)
	}

	n := len(arr)
	if n == 0 {
		fmt.Println("Tidak ada data yang dimasukkan.")
		return
	}

	insertionSort(arr, n)

	isConstant, difference := isConstantDifference(arr, n)

	fmt.Println("Array setelah diurutkan:")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	if isConstant {
		fmt.Printf("Data berjarak tetap dengan selisih %d\n", difference)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
