package main

import "fmt"

const maxSize = 1000000

func sortArray(data2217 []int, ukuran int) {
	for start2 := 0; start2 < ukuran-1; start2++ {
		smallest := start2
		for check := start2 + 1; check < ukuran; check++ {
			if data2217[check] < data2217[smallest] {
				smallest = check
			}
		}
		data2217[start2], data2217[smallest] = data2217[smallest], data2217[start2]
	}
}

func hitungMedian217(data []int, size int) float64 {
	mid := size / 2
	if size%2 == 0 {
		return float64(data[mid-1]+data[mid]) / 2.0
	}
	return float64(data[mid])
}

func main() {
	var values [maxSize]int
	var sorted [maxSize]int
	var count int
	var input int

	fmt.Println("Masukkan bilangan bulat (gunakan 0 untuk menampilkan median sementara, akhiri dengan -5313):")
	for {
		fmt.Scan(&input)
		if input == -5313 {
			break
		}

		if input != 0 {
			values[count] = input
			count++
		} else if count > 0 {
			for i := 0; i < count; i++ {
				sorted[i] = values[i]
			}

			sortArray(sorted[:], count)
			median := hitungMedian217(sorted[:], count)
			fmt.Printf("%.0f\n", median)
		}
	}
}
