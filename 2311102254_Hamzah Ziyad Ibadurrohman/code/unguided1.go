package main

import "fmt"

func selectionSort(arr []int, ascending bool) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		extremeIdx := i
		for j := i + 1; j < n; j++ {
			if (ascending && arr[j] < arr[extremeIdx]) || (!ascending && arr[j] > arr[extremeIdx]) {
				extremeIdx = j
			}
		}
		arr[i], arr[extremeIdx] = arr[extremeIdx], arr[i]
	}
	return arr
}
func prosesrumah(rumah []int) {
	var angkaganjil, angkagenap []int

	// Memisahkan bilangan ganjil dan genap
	for _, num := range rumah {
		if num%2 == 0 {
			angkagenap = append(angkagenap, num)
		} else {
			angkaganjil = append(angkaganjil, num)
		}
	}

	// Urutkan bilangan ganjil (membesar) dan genap (mengecil)
	angkaganjil = selectionSort(angkaganjil, true)
	angkagenap = selectionSort(angkagenap, false)

	// Cetak bilangan ganjil diikuti bilangan genap
	for _, num := range angkaganjil {
		fmt.Printf("%d ", num)
	}
	for _, num := range angkagenap {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

}
func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		rumah := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&rumah[j])
		}
		prosesrumah(rumah)
	}

}
