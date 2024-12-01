package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	var data []int

	fmt.Println("Masukkan angka satu per satu. Gunakan 0 untuk mencetak median, dan akhiri dengan -5313 untuk keluar.")

	for {
		var num int
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&num)

		if num == -5313 {
			fmt.Println("Program selesai.")
			break
		}

		if num == 0 {
			if len(data) == 0 {
				fmt.Println("Tidak ada data untuk menghitung median.")
				continue
			}

			selectionSort(data)

			length := len(data)
			if length%2 == 0 {
				mid1 := data[length/2-1]
				mid2 := data[length/2]
				fmt.Printf("Median: %d\n", (mid1+mid2)/2)
			} else {
				fmt.Printf("Median: %d\n", data[length/2])
			}
			data = nil
		} else {
			data = append(data, num)
		}
	}
}
