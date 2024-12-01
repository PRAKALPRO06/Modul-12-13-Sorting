package main

import (
	"fmt"
)

func pengurutanNomorRumah(arr []int, pengurutanNaik bool) {
	x := len(arr)
	for i := 0; i < x-1; i++ {
		idx := i
		for j := i + 1; j < x; j++ {
			if (pengurutanNaik && arr[j] < arr[idx]) || (!pengurutanNaik && arr[j] > arr[idx]) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func main() {
	var x int
	fmt.Print("Masukkan jumlah daerah kerabat (x): ")
	if _, err := fmt.Scan(&x); err != nil || x <= 0 {
		fmt.Println("Input jumlah daerah harus berupa angka lebih dari 0.")
		return
	}

	for daerah := 1; daerah <= x; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		if _, err := fmt.Scan(&m); err != nil || m <= 0 {
			fmt.Println("Input jumlah nomor rumah harus berupa angka lebih dari 0.")
			return
		}

		arr := make([]int, m)
		ganjil := []int{}
		genap := []int{}
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			if _, err := fmt.Scan(&arr[i]); err != nil {
				fmt.Println("Input nomor rumah tidak valid. Harus berupa angka.")
				return
			}
			if arr[i]%2 == 0 {
				genap = append(genap, arr[i])
			} else {
				ganjil = append(ganjil, arr[i])
			}
		}

		pengurutanNomorRumah(ganjil, true)
		pengurutanNomorRumah(genap, false)

		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
