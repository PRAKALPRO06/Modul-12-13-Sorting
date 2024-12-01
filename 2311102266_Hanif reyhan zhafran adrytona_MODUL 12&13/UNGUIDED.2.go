// 2311102266_Hanif Reyhan Zhafran Arytona

package main

import (
	"fmt"
	"sort"
)

func calculateMedian(arr []int) float64 {
	n := len(arr)
	if n%2 == 0 {
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	}
	return float64(arr[n/2])
}

func main() {
	var data []int
	var num int

	fmt.Println("Masukkan angka yang dipisahkan dengan spasi (akhiri dengan -5313):")

	for {
		fmt.Scan(&num)
		if num == -5313 {
			break
		}
		if num == 0 {
			sort.Ints(data)
			median := calculateMedian(data)
			fmt.Printf("Median: %.0f\n", median)
		} else {
			data = append(data, num)
		}
	}
}
