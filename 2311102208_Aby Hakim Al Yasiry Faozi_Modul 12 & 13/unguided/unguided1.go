package main

import (
	"fmt"
)

func ganjilGenap(arr []int) (ganjil []int, genap []int) {
	for _, nomor := range arr {
		if nomor%2 == 0 {
			genap = append(genap, nomor)
		} else {
			ganjil = append(ganjil, nomor)
		}
	}
	return
}

func selectionSort(arr []int, naik bool) {
	panjang := len(arr)
	for i := 0; i < panjang-1; i++ {
		posisi := i
		for j := i + 1; j < panjang; j++ {
			if (naik && arr[j] < arr[posisi]) || (!naik && arr[j] > arr[posisi]) {
				posisi = j
			}
		}
		arr[i], arr[posisi] = arr[posisi], arr[i]
	}
}

func main() {
	var jumlahDaerah int
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&jumlahDaerah)

	for i := 0; i < jumlahDaerah; i++ {
		fmt.Printf("\nMasukkan jumlah rumah di daerah ke-%d: ", i+1)
		var jumlahRumah int
		fmt.Scan(&jumlahRumah)

		nomorRumah := make([]int, jumlahRumah)
		fmt.Printf("Masukkan nomor rumah di daerah ke-%d:\n", i+1)
		for j := 0; j < jumlahRumah; j++ {
			fmt.Scan(&nomorRumah[j])
		}

		ganjil, genap := ganjilGenap(nomorRumah)

		selectionSort(ganjil, false)
		selectionSort(genap, true)

		fmt.Printf("\nHasil untuk daerah ke-%d:\n", i+1)
		for _, nomor := range ganjil {
			fmt.Printf("%d ", nomor)
		}
		for _, nomor := range genap {
			fmt.Printf("%d ", nomor)
		}
		fmt.Println()
	}
}
