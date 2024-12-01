package main

import (
	"fmt"
	"sort"
)


func hitungMedian(data []int) float64 {
	jumlah := len(data)
	if jumlah%2 == 1 {
		return float64(data[jumlah/2]) 
	}

	return float64(data[jumlah/2-1]+data[jumlah/2]) / 2.0
}

func main() {
	var data []int
	var angka int

	fmt.Println("Masukkan data bilangan bulat (akhiri dengan -5313):")
	for {
		fmt.Scan(&angka)
		if angka == -5313 {
			break
		}
		if angka != 0 {
			data = append(data, angka)
		} else {
			if len(data) == 0 {
				fmt.Println("Median: Tidak ada data.")
			} else {
				sort.Ints(data)
				median := hitungMedian(data)
				fmt.Printf("Median: %.0f\n", median)
			}
			data = []int{}
		}
	}
}
