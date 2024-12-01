//2311102010_Rakha Yudhistira_IF-11-06

package main

import "fmt"

// Fungsi untuk mengurutkan array menggunakan Selection Sort
func selectionSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen minimum dengan elemen pertama
		numbers[i], numbers[minIdx] = numbers[minIdx], numbers[i]
	}
}

// Fungsi untuk menghitung median
func calculateMedian(numbers []int) int {
	selectionSort(numbers) // Urutkan data menggunakan Selection Sort
	n := len(numbers)

	if n%2 == 1 {
		return numbers[n/2]
	}
	return (numbers[n/2-1] + numbers[n/2]) / 2
}

func main() {
	var input int
	var numbers []int

	fmt.Println("Masukkan angka satu per satu, akhiri dengan -5313:")

	for {
		fmt.Scan(&input)

		// Akhiri jika menemukan -5313
		if input == -5313 {
			break
		}

		// Abaikan angka 0
		if input != 0 {
			numbers = append(numbers, input)
		}
	}

	if len(numbers) == 0 {
		fmt.Println("Tidak ada angka untuk diproses.")
		return
	}

	// Hitung median
	median := calculateMedian(numbers)
	fmt.Println("Median:", median)
}
