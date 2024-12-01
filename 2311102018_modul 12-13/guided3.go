//2311102018 - Modul 12 Aryo Tegar Sukarno
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var numbers []int
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan angka (akhiri dengan -5313):")

	for {
		// Membaca input dari pengguna
		reader.Scan()
		input := reader.Text()

		// Konversi input menjadi integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan angka!")
			continue
		}

		// Berhenti membaca jika angka -5313 ditemukan
		if num == -5313 {
			break
		}

		// Tambahkan angka ke dalam slice
		numbers = append(numbers, num)

		// Mengurutkan slice
		sort.Ints(numbers)

		// Hitung median
		median := calculateMedian(numbers)

		// Cetak median saat ini
		fmt.Printf("Data tersimpan: %v\n", numbers)
		fmt.Printf("Median saat ini: %d\n", median)
	}

	fmt.Println("Proses selesai. Terima kasih!")
}

// Fungsi untuk menghitung median
func calculateMedian(numbers []int) int {
	n := len(numbers)
	if n%2 == 0 {
		// Jika jumlah elemen genap, median adalah rata-rata dari dua nilai tengah
		return (numbers[n/2-1] + numbers[n/2]) / 2
	}
	// Jika jumlah elemen ganjil, median adalah elemen tengah
	return numbers[n/2]
}
