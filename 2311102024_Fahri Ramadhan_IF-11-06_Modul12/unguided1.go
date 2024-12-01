package main

import (
	"fmt"
	"sort"
)

func pisahkanGanjilGenap(arr []int) ([]int, []int) {
	var ganjil, genap []int
	for _, num := range arr {
		if num%2 == 1 {
			ganjil = append(ganjil, num)
		} else {
			genap = append(genap, num)
		}
	}
	return ganjil, genap
}

func urutkanNomorRumah(ganjil []int, genap []int) []int {
	sort.Ints(ganjil)
	sort.Sort(sort.Reverse(sort.IntSlice(genap)))
	return append(ganjil, genap...)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)
	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		ganjil, genap := pisahkanGanjilGenap(arr)
		hasil := urutkanNomorRumah(ganjil, genap)

		fmt.Printf("Nomor rumah terurut untuk dikunjungi di daerah %d: ", daerah)
		for _, num := range hasil {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}