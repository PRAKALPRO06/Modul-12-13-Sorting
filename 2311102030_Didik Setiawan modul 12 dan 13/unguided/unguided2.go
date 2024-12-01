package main

import (
	"fmt"
	"sort"
)

func median(arr []int) int {
	n := len(arr)
	if n%2 == 0 {
		return (arr[n/2-1] + arr[n/2]) / 2
	}
	return arr[n/2]
}

func main() {
	var arr []int
	var arrMedian []int

	fmt.Println("Masukkan data (akhiri dengan -5313): ")
	for {
		var datum int
		fmt.Scan(&datum)
		if datum == -5313 {
			break
		}
		if datum == 0 {
			sort.Ints(arr) // Sort array
			if len(arr) > 0 {
				arrMedian = append(arrMedian, median(arr))
			}
		} else {
			arr = append(arr, datum)
		}
	}

	fmt.Println("Output:")
	for _, med := range arrMedian {
		fmt.Println(med)
	}
}
