package main

import (
	"fmt"
	"sort"
)

func main() {
	var number int
	var numbers []int

	fmt.Println("Masukkan bilangan (akhiri dengan -5313):")
	for {
		fmt.Scan(&number)

		if number == -5313 {
			break
		}

		if number == 0 {
			if len(numbers) > 0 {
				sort.Ints(numbers)
				median := findMedian(numbers)
				fmt.Printf("Median: %d\n", median)
			}
		} else {
			numbers = append(numbers, number)
		}
	}
}

func findMedian(numbers []int) int {
	length := len(numbers)
	if length%2 == 1 {
		return numbers[length/2]
	}
	return (numbers[(length/2)-1] + numbers[length/2]) / 2
}