package main

import "fmt"

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func cekJarak(data []int) string {
	if len(data) < 2 {
		return "Data berjarak tidak tetap"
	}
	distance := data[1] - data[0]
	for i := 2; i < len(data); i++ {
		if data[i]-data[i-1] != distance {
			return "Data berjarak tidak tetap"
		}
	}
	return fmt.Sprintf("Data berjarak %d", distance)
}

func main() {
	var data []int
	var input int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	insertionSort(data)

	for i, num := range data {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()

	fmt.Println(cekJarak(data))
}
