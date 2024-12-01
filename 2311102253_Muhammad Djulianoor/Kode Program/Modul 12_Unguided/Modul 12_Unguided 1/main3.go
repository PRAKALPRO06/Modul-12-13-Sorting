package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSort(arr_2311102253 []int, menaik bool) {
	n := len(arr_2311102253)
	for i := 0; i < n-1; i++ {
		extremeIdx := i
		for j := i + 1; j < n; j++ {
			if menaik {
				if arr_2311102253[j] < arr_2311102253[extremeIdx] {
					extremeIdx = j
				}
			} else {
				if arr_2311102253[j] > arr_2311102253[extremeIdx] {
					extremeIdx = j
				}
			}
		}
		arr_2311102253[i], arr_2311102253[extremeIdx] = arr_2311102253[extremeIdx], arr_2311102253[i]
	}
}

func processRow(row string) string {
	x := strings.Fields(row)

	var odds, evens []int

	for _, numStr := range x {
		num, _ := strconv.Atoi(numStr)
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}

	selectionSort(odds, false)

	selectionSort(evens, true)

	var result strings.Builder
	for _, odd := range odds {
		result.WriteString(fmt.Sprintf("%d ", odd))
	}
	for _, even := range evens {
		result.WriteString(fmt.Sprintf("%d ", even))
	}

	return strings.TrimSpace(result.String())
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jumlah baris: ")
	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, _ := strconv.Atoi(nStr)

	baris := make([]string, n)
	fmt.Println("Masukkan data baris: ")
	for i := 0; i < n; i++ {
		baris[i], _ = reader.ReadString('\n')
		baris[i] = strings.TrimSpace(baris[i])
	}

	fmt.Println("Keluaran:")
	for i, row := range baris {
		fmt.Printf("%d %s\n", i+1, processRow(row))
	}
}
