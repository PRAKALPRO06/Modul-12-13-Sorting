package main

import "fmt"

func printGanjilGenap(arr []int, n int)  {
	arrGanjil := make([]int, n)
	arrGenap := make([]int, n)
	idxGenap := 0
	idxGanjil := 0

	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			arrGenap[idxGenap] = arr[i]
			idxGenap++
		} else {
			arrGanjil[idxGanjil] = arr[i]
			idxGanjil++
		}
		
	}

	selectionSort(arrGenap, arrGanjil, idxGenap, idxGanjil)

	for i := 0; i < idxGanjil; i++ {
		fmt.Print(arrGanjil[i], " ")
	}
	for i := 0; i < idxGenap; i++ {
		fmt.Print(arrGenap[i], " ")
	}

}

func selectionSort(arrGenap, arrGanjil []int, idxGenap, idxGanjil int) {
	var arrass int
	var arrdess int

	for i := 0; i < idxGanjil; i++ {
		for j := i+1; j < idxGanjil; j++ {
			if arrGanjil[i] > arrGanjil[j] {
				arrass = arrGanjil[j]
				arrGanjil[j] = arrGanjil[i]
				arrGanjil[i] = arrass
			}
		}
	}
	for i := 0; i < idxGenap; i++ {
		for j := i+1; j < idxGenap; j++ {
			if arrGenap[i] < arrGenap[j] {
				arrdess = arrGenap[j]
				arrGenap[j] = arrGenap[i]
				arrGenap[i] = arrdess
			}
		}
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	// Proses tiap daerah
	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		// Membaca nomor rumah untuk daerah ini
		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		
		// Urutkan array dari terkecil ke terbesar
		// Tampilkan hasil
		printGanjilGenap(arr, m)

		fmt.Println()
	}
}