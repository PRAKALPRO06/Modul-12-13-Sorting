// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import "fmt"

// Fungsi untuk mengurutkan array dengan selection sort
func selectionSort(arr []int, asc bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (asc && arr[j] < arr[idx]) || (!asc && arr[j] > arr[idx]) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

// Fungsi untuk mencetak garis batas panjang
func printSeparator() {
	fmt.Println("\n=================================================")
}

func printDashedSeparator() {
	fmt.Println("-------------------------------------------------")
}

// Fungsi untuk menampilkan nomor rumah berdasarkan daerah
func processDaerah(i, m int) ([]int, []int) {
	var arr = make([]int, m)
	// Input nomor rumah kerabat untuk setiap daerah
	fmt.Printf("Masukkan %d nomor rumah kerabat untuk daerah ke-%d: ", m, i+1)
	for j := 0; j < m; j++ {
		fmt.Scan(&arr[j])
	}

	// Memisahkan nomor ganjil dan genap
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

// Fungsi untuk menampilkan hasil terurut
func printSortedResults(i int, odd, even []int) {
	// Mengurutkan nomor rumah
	selectionSort(odd, true)   // urutkan ganjil secara menaik
	selectionSort(even, false) // urutkan genap secara menurun

	// Output nomor rumah yang terurut
	fmt.Printf("\nNomor rumah terurut untuk daerah ke-%d:\n", i+1)

	// Tampilkan nomor ganjil
	for _, num := range odd {
		fmt.Printf("%d ", num)
	}

	// Tampilkan nomor genap
	for _, num := range even {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// Fungsi utama program
func main() {
	var n int
	// Input jumlah daerah kerabat
	printSeparator()
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	// Proses untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		// Input jumlah rumah kerabat untuk setiap daerah
		printDashedSeparator()
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah ke-%d (m): ", i+1)
		fmt.Scan(&m)

		// Ambil nomor rumah untuk daerah tersebut
		odd, even := processDaerah(i, m)

		// Menampilkan hasil terurut untuk daerah tersebut
		printSortedResults(i, odd, even)

		// Cetak garis pemisah untuk setiap daerah
		printDashedSeparator()
	}
}
