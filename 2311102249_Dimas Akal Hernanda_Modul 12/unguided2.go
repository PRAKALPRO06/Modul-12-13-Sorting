package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Fungsi untuk menghitung median dari array yang sudah diurutkan
func calculateMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 {
		// Jika jumlah data ganjil, ambil nilai tengah
		return arr[n/2]
	} else {
		// Jika jumlah data genap, hitung rata-rata dua nilai tengah
		return (arr[n/2-1] + arr[n/2]) / 2
	}
}

func main() {
	var data []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan bilangan bulat:")

	// Membaca input dari pengguna
	if scanner.Scan() {
		input := scanner.Text()
		strNums := strings.Fields(input)

		for _, str := range strNums {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Masukkan hanya bilangan bulat!")
				return
			}

			if num == -5313 {
				// Marker akhir input, keluar dari loop
				break
			} else if num == 0 {
				// Ketika menemukan angka 0, urutkan data dan cetak median
				sort.Ints(data) // Mengurutkan data
				if len(data) > 0 {
					median := calculateMedian(data)
					fmt.Println("Median:", median)
				} else {
					fmt.Println("Tidak ada data untuk menghitung median.")
				}
			} else if num > 0 {
				// Tambahkan bilangan positif ke array
				data = append(data, num)
			}
		}
	}
}
