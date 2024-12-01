package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Masukkan :")
	fmt.Println("==============================================")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	masukan := scanner.Text()

	angkaInput := strings.Split(masukan, " ")
	fmt.Println("==============================================")
	fmt.Printf("\n")

	var angka []int
	fmt.Println("Keluaran :")
	fmt.Println("==============================================")
	for _, bilanganStr := range angkaInput {
		bilangan, _ := strconv.Atoi(bilanganStr)

		if bilangan == -5313 {
			break
		}

		if bilangan == 0 {
			median := hitungMedian(angka)
			fmt.Println(median)
		} else {
			angka = append(angka, bilangan)
		}
	}
	fmt.Println("==============================================")
}

func hitungMedian(angka []int) int {
	sort.Ints(angka)

	jumlah := len(angka)
	if jumlah%2 == 1 {
		return angka[jumlah/2]
	} else {
		return (angka[jumlah/2-1] + angka[jumlah/2]) / 2
	}
}