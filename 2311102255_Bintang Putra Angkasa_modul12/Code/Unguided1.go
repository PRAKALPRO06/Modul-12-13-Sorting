package main

import "fmt"

const maxArraySize = 1000000

func calculateMedian(data []int) int {
	n := len(data)
	if n%2 == 1 {
		return data[n/2]
	}
	return (data[n/2-1] + data[n/2]) / 2
}

func insertionSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
	return data
}

func main() {
	var data []int      
	var result []int    
	var num int

	fmt.Println("Masukkan data (akhiri dengan -5313):")

	for {
		fmt.Scan(&num)

		if num == -5313 {
			for _, r := range result {
				fmt.Println(r)
			}
			return
		} else if num == 0 {	
			data = insertionSort(data) 
			median := calculateMedian(data)
			result = append(result, median)
		} else {
			
			if len(data) >= maxArraySize {
				fmt.Println("Error: Jumlah data melebihi batas maksimum 1.000.000 elemen.")
				continue
			}
			
			data = append(data, num)
		}
	}
}
