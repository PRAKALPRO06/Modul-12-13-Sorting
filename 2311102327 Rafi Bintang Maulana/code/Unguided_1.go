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
	fmt.Println("Masukan")
	fmt.Println("================================")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	jumlahDaerah, _ := strconv.Atoi(scanner.Text())

	var hasil [][]int

	for i := 0; i < jumlahDaerah; i++ {
		scanner.Scan()
		baris := scanner.Text()
		data := strings.Fields(baris)

		var ganjil, genap []int
		for _, nomorStr := range data[1:] {
			nomor, _ := strconv.Atoi(nomorStr)
			if nomor%2 == 0 {
				genap = append(genap, nomor)
			} else {
				ganjil = append(ganjil, nomor)
			}
		}

		sort.Ints(genap)
		sort.Sort(sort.Reverse(sort.IntSlice(ganjil)))

		hasil = append(hasil, append(ganjil, genap...))
	}
	fmt.Println("================================")

	fmt.Printf("\n")
	fmt.Println("Keluaran")
	fmt.Println("================================")

	for _, daerah := range hasil {
		for i, nomor := range daerah {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(nomor)
		}
		fmt.Println()
	}
	fmt.Printf("================================ \n")
}