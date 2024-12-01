package main

import "fmt"

func median(data []int) int {
	n := len(data)
	if n%2 == 1 {
		// Jika jumlah data ganjil, ambil elemen tengah
		return data[n/2]
	} else {
		// Jika jumlah data genap, ambil rata-rata dua elemen tengah dan dibulatkan ke bawah
		return (data[n/2-1] + data[n/2]) / 2
	}
}

// Fungsi Insertion Sort
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

func main() {
	var bilangan int
	var data []int

	for {
		fmt.Scan(&bilangan)

		if bilangan == -5313 {
			break
		} else if bilangan == 0 {
			insertionSort(data)
			fmt.Println(median(data))
		} else {
			data = append(data, bilangan)
		}
	}
}
