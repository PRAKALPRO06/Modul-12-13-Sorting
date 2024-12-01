package main

import (
	"fmt"
	"sort"
)

func processDaerah(m int) ([]int, []int) {
	arr := make([]int, m)
	fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
	for j := 0; j < m; j++ {
		fmt.Scan(&arr[j])
	}

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

func printSortedResults(i int, odd, even []int) {
	sort.Ints(odd)                               // Sort ascending
	sort.Sort(sort.Reverse(sort.IntSlice(even))) // Sort descending

	fmt.Printf("\nNomor rumah terurut untuk daerah ke-%d:\n", i+1)
	fmt.Println("Odd numbers: ", odd)
	fmt.Println("Even numbers:", even)
}

func main() {
	var n int
	fmt.Println("\n=================================================")
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Println("\n-------------------------------------------------")
		fmt.Printf("Masukkan jumlah rumah kerabat untuk daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		odd, even := processDaerah(m)
		printSortedResults(i, odd, even)
	}

	fmt.Println("\nProgram selesai.")
}
