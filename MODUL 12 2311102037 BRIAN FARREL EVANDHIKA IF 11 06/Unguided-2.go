// 2311102037_BRIAN FARREL EVANDHIKA_IF 11_06
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk mengurutkan array menggunakan algoritma Selection Sort
func sortArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// Fungsi untuk menghitung median dari array yang telah diurutkan
func getMedian(arr []int) float64 {
	length := len(arr)
	if length%2 == 0 {
		return float64(arr[length/2-1]+arr[length/2]) / 2.0
	}
	return float64(arr[length/2])
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan angka-angka yang dipisahkan dengan spasi (akhiri dengan -5313):")

	reader.Scan()
	input := reader.Text()
	values := strings.Split(input, " ")
	var numbers []int

	for _, val := range values {
		number, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan hanya angka bulat.")
			return
		}

		if number == -5313 {
			break
		} else if number == 0 {
			sortArray(numbers)
			median := getMedian(numbers)
			fmt.Printf("%.0f\n", median)
		} else {
			numbers = append(numbers, number)
		}
	}
}
