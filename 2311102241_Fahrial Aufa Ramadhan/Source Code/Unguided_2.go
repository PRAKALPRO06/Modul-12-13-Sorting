package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&n)

	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("Masukkan jumlah nomor %d: ", daerah)
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		ganjil := []int{}
		genap := []int{}
		for _, num := range arr {
			if num%2 == 1 {
				ganjil = append(ganjil, num)
			} else {
				genap = append(genap, num)
			}
		}

		sort.Ints(ganjil)
		sort.Sort(sort.Reverse(sort.IntSlice(genap)))

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
