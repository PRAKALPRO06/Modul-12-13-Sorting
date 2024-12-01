package main

import "fmt"

//Erwin Rivaldo Silaban/2311102248

func selectionSort(arr []int, asc bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			// Kondisi pengurutan berdasarkan parameter asc
			if (asc && arr[j] < arr[idx]) || (!asc && arr[j] > arr[idx]) {
				idx = j
			}
		}
		// Tukar elemen jika diperlukan
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

// processDaerah memproses rumah dalam satu daerah
func processDaerah(i, m int) ([]int, []int) {
	// Buat slice untuk menampung nomor rumah
	var arr = make([]int, m)

	// Input nomor rumah
	fmt.Printf("Masukkan %d nomor rumah kerabat untuk daerah ke-%d: ", m, i+1)
	for j := 0; j < m; j++ {
		fmt.Scan(&arr[j])
	}

	// Pisahkan nomor rumah genap dan ganjil
	var odd, even []int
	for _, num := range arr {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return odd, even
}

// printSortedResults mencetak hasil pengurutan
func printSortedResults(i int, odd, even []int) {
	selectionSort(odd, true)
	selectionSort(even, false)

	// Cetak header
	fmt.Printf("\nNomor Rumah Setelah di Urutkan di daerah ke-%d:\n", i+1)

	// Cetak nomor ganjil
	for _, num := range odd {
		fmt.Printf("%d ", num)
	}
	// Cetak nomor genap
	for _, num := range even {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func main() {
	// Variabel untuk jumlah daerah
	var n int
	fmt.Print("Masukkan Jumlah Daerah : ")
	fmt.Scan(&n)

	// Proses setiap daerah
	for i := 0; i < n; i++ {
		// Variabel untuk jumlah rumah di setiap daerah
		var m int
		fmt.Printf("Masukkan Jumlah Rumah Kerabat di Daerah ke-%d : ", i+1)
		fmt.Scan(&m)
		odd, even := processDaerah(i, m)
		printSortedResults(i, odd, even)
	}
}
