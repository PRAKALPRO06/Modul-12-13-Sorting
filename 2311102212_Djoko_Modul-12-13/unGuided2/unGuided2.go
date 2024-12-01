package main

import "fmt"

// Fungsi untuk mengurutkan data menggunakan selection sort
func selectionSort(data []int) {
	n := len(data)
	i := 1
	for i <= n-1 {
		idxMin := i - 1
		j := i
		for j < n {
			if data[idxMin] > data[j] {
				idxMin = j
			}
			j = j + 1
		}
		// Tukar elemen
		t := data[idxMin]
		data[idxMin] = data[i-1]
		data[i-1] = t
		i = i + 1
	}
}

// Fungsi untuk menghitung median
func calculateMedian(data []int) int {
	// Mengurutkan data dengan selection sort
	selectionSort(data)

	// Menentukan median
	n := len(data)
	if n%2 == 1 {
		return data[n/2] // Jika ganjil, median adalah elemen tengah
	} else {
		return (data[n/2-1] + data[n/2]) / 2 // Jika genap, median adalah rata-rata dua elemen tengah
	}
}

func main() {
	var input2311102212 int
	var data []int
	var i int
	i = 1
	for {
		fmt.Printf("Masukan data ke-%d = ", i)
		fmt.Scan(&input2311102212)

		// Jika input adalah -5313, hentikan program
		if input2311102212 == -5313 {
			break
		}

		// Jika input adalah 0, hitung dan cetak median
		if input2311102212 == 0 {
			if len(data) == 0 {
				fmt.Println("Tidak ada data")
			} else {
				median := calculateMedian(data)
				fmt.Println(median)
			}
		} else {
			// Masukkan angka ke dalam data (selain 0 dan -5313)
			data = append(data, input2311102212)
		}
		i++
	}
}
