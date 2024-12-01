package main

import (
	"fmt"
)

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

func CariMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	} else {
		return (arr[n/2-1] + arr[n/2]) / 2
	}
}

func main() {
	var data []int
	for {
		var num int
		fmt.Scan(&num)

		if num == -5313 {
			break
		}
		if num == 0 {
			insertionSort(data)
			median := CariMedian(data)
			fmt.Println(median)
		} else {
			data = append(data, num)
		}
	}
}
