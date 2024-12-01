package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func hitungJarak(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0
	}
	spacing := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != spacing {
			return false, 0
		}
	}
	return true, spacing
}

func main() {
	var data []int
	var input int

	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Terjadi kesalahan input:", err)
			return
		}
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	if len(data) == 0 {
		fmt.Println("Tidak ada data yang diproses.")
		return
	}

	insertionSort(data)

	fmt.Println("Data setelah diurutkan:")
	for _, num := range data {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	isConstant, spacing := hitungJarak(data)
	if isConstant {
		fmt.Printf("Data berjarak %d\n", spacing)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
