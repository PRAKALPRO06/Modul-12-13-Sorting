package main

import (
	"fmt"
	"sort"
)

func pisahganjilgenap_2311102033(arr []int) ([]int, []int) {
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

func urutnomerrumah_2311102033(ganjil []int, genap []int) []int {
	sort.Ints(ganjil)
	sort.Sort(sort.Reverse(sort.IntSlice(genap)))
	return append(ganjil, genap...)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	// Garis pemisah
	fmt.Println("\n" + string('-'*40))

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("Masukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		ganjil, genap := pisahganjilgenap_2311102033(arr)
		hasil := urutnomerrumah_2311102033(ganjil, genap)

		// Garis pemisah sebelum hasil
		fmt.Println("\n" + string('='*40))
		fmt.Printf("Nomor rumah terurut untuk dikunjungi di daerah %d: ", daerah)
		for _, num := range hasil {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
		// Garis pemisah setelah hasil
		fmt.Println(string('-' * 40))
	}
}
