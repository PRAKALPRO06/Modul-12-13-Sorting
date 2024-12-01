package main

import (
	"fmt"
	"sort"
)

func pemisahGanjilGenap(arr []int) ([]int, []int) {
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

func urutanNomor(ganjil []int, genap []int) []int {
	sort.Ints(ganjil)
	sort.Sort(sort.Reverse(sort.IntSlice(genap)))
	return append(ganjil, genap...)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&n)
	fmt.Println("\n" + string('-'*40))

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("Masukkan jumlah nomor rumah %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		ganjil, genap := pemisahGanjilGenap(arr)
		hasil := urutanNomor(ganjil, genap)

		fmt.Println("\n" + string('='*40))
		fmt.Printf("Nomor terurut %d: ", daerah)
		for _, num := range hasil {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
		fmt.Println(string('-' * 40))
	}
}
