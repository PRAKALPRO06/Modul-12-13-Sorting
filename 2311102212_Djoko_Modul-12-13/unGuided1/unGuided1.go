package main

import "fmt"

// Fungsi untuk mengurutkan array menggunakan Selection Sort
func selectionSort(arr []int, ascending bool) {
	n := len(arr)
	i := 1
	for i <= n-1 {
		idxMin := i - 1
		j := i
		for j < n {
			// Kondisi untuk ascending atau descending
			if (ascending && arr[idxMin] < arr[j]) || (!ascending && arr[idxMin] > arr[j]) {
				idxMin = j
			}
			j = j + 1
		}
		// Tukar elemen
		t := arr[idxMin]
		arr[idxMin] = arr[i-1]
		arr[i-1] = t
		i = i + 1
	}
}

// Fungsi untuk memisahkan nomor ganjil dan genap, serta mengurutkan secara manual
func processRumahManual(arr []int) []int {
	var ganjil, genap []int

	// Memisahkan angka ganjil dan genap
	for _, num := range arr {
		if num%2 == 0 {
			genap = append(genap, num)
		} else {
			ganjil = append(ganjil, num)
		}
	}

	// Mengurutkan ganjil secara menurun (descending) dengan selection sort
	selectionSort(ganjil, false)

	// Mengurutkan genap secara menaik (ascending) dengan selection sort
	selectionSort(genap, true)

	// Menggabungkan ganjil dan genap
	return append(ganjil, genap...)
}

func main() {
	var banyakDaerah2311102212 int
	fmt.Print("Masukan Berapa Daerah: ")
	fmt.Scan(&banyakDaerah2311102212)

	// Menggunakan slice dua dimensi untuk menyimpan data rumah
	rumah := make([][]int, banyakDaerah2311102212)

	for i := 0; i < banyakDaerah2311102212; i++ {
		var banyakRumah int
		fmt.Printf("Masukan Banyak Rumah di Daerah %d: ", i+1)
		fmt.Scan(&banyakRumah)

		// Membuat slice untuk menyimpan nomor rumah di daerah ini
		rumah[i] = make([]int, banyakRumah)
		for j := 0; j < banyakRumah; j++ {
			fmt.Printf("Masukan Rumah Nomor Ke-%d: ", j+1)
			fmt.Scan(&rumah[i][j])
		}
	}

	// Memproses dan menampilkan hasil
	fmt.Println("\nHasil:")
	for i, daerah := range rumah {
		hasil := processRumahManual(daerah)
		fmt.Printf("Daerah %d: ", i+1)
		for _, num := range hasil {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
