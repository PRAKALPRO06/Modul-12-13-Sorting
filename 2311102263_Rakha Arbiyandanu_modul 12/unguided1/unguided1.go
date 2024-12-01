package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Print("Masukkan jumlah daerah kerabat: ")
	fmt.Scan(&t)

	for daerah := 1; daerah <= t; daerah++ {
		var n int
		fmt.Printf("Masukkan jumlah nomor rumah untuk daerah %d: ", daerah)
		fmt.Scan(&n)

		arr := make([]int, n)
		fmt.Printf("Masukkan %d nomor rumah: ", n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		ganjil := []int{}
		genap := []int{}
		for _, num := range arr {
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}

		sort.Slice(ganjil, func(i, j int) bool {
			return ganjil[i] > ganjil[j]
		})

		sort.Ints(genap)

		// Cetak hasil
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
