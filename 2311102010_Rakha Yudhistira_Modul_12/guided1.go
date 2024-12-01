//2311102010_Rakha Yudhistira_IF-11-06



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

func main() {
	var n int
	fmt.Scan(&n)

	var results [][]int

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah)

		results = append(results, rumah)
	}

	for _, rumah := range results {
		for j := 0; j < len(rumah); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[j])
		}
		fmt.Println()
	}
}