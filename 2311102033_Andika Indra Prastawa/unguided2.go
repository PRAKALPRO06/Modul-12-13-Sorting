package main

import (
	"fmt"
	"sort"
)

func main() {
	var input_2311102033 int
	var data_2311102033 []int

	fmt.Println("Masukkan bilangan (Dan akhiri dengan -5313):")
	for {
		fmt.Scan(&input_2311102033)

		if input_2311102033 == -5313 {
			break
		}

		if input_2311102033 == 0 {
			if len(data_2311102033) > 0 {
				sort.Ints(data_2311102033)
				median := menghitungMedian_2311102033(data_2311102033)
				fmt.Printf(": %d\n", median)
			}
		} else {
			data_2311102033 = append(data_2311102033, input_2311102033)
		}
	}
}

func menghitungMedian_2311102033(data_2311102033 []int) int {
	n := len(data_2311102033)
	if n%2 == 1 {
		return data_2311102033[n/2]
	}
	return (data_2311102033[(n/2)-1] + data_2311102033[n/2]) / 2
}
