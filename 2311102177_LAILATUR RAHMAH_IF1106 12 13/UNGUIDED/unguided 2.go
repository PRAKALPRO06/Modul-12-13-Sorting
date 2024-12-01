// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk melakukan selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Cari elemen terkecil di bagian yang belum terurut
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen terkecil yang ditemukan dengan elemen pertama
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Fungsi untuk menghitung median dari array yang sudah terurut
func calculateMedian(arr []int) float64 {
	n := len(arr)
	if n%2 == 0 {
		// Jika jumlah elemen genap, median adalah rata-rata dari dua nilai tengah
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	}
	// Jika jumlah elemen ganjil, median adalah elemen tengah
	return float64(arr[n/2])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Membaca seluruh input sebagai satu baris
	fmt.Println("Masukkan angka yang dipisahkan dengan spasi (akhiri dengan -5313):")
	scanner.Scan()
	line := scanner.Text()

	// Memisahkan input ke dalam bagian-bagian
	parts := strings.Split(line, " ")
	var data []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Input tidak valid, harap masukkan angka bulat saja.")
			return
		}

		if num == -5313 {
			// Akhir dari input
			break
		} else if num == 0 {
			// Urutkan data dan hitung median
			selectionSort(data)
			median := calculateMedian(data)
			fmt.Printf("%.0f\n", median)
		} else {
			// Tambahkan angka ke array data
			data = append(data, num)
		}
	}
}
