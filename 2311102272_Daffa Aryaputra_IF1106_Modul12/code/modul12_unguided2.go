package main

import (
	"fmt"
	"sort"
)

func main() {
	var input int
	var data []int

	fmt.Println("Masukkan bilangan (akhiri dengan -5313):")
	for {
		fmt.Scan(&input)

		if input == -5313 {
			break
		}

		if input == 0 {
			if len(data) > 0 {
				sort.Ints(data)
				median := calculateMedian(data)
				fmt.Printf("Median: %d\n", median)
			}
		} else {
			data = append(data, input)
		}
	}
}

func calculateMedian(data []int) int {
	n := len(data)
	if n%2 == 1 {
		return data[n/2]
	}
	return (data[(n/2)-1] + data[n/2]) / 2
}
